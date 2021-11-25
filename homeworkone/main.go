package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	_ "net/http/pprof"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", rootHandler)
	mux.HandleFunc("/healthz", healthz)
	var PORT string
	var ADDRESS string
	PORT = os.Getenv("service_port")
	ADDRESS = os.Getenv("service_address")
	log.Println(PORT)
	serv := &http.Server{
		Addr:    ADDRESS+":"+PORT,
		Handler: mux,
	}
	log.Println(serv.Addr)
	go func() {
		if err := serv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("start http server failed: %s\n", err)
		}
	}()
	// grace shutdown
	quit := make(chan os.Signal, 1)
	// receive system signal
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM,syscall.SIGQUIT)
	<-quit // block
	// service will be shut down in 5 seconds, wait for the request to be processed
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := serv.Shutdown(ctx); err != nil {
		log.Fatalf("shutdown server failed: %s", err)
	}
	log.Println("server shutdown successfully")
}

func healthz(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	var VERSION string
	VERSION = os.Getenv("VERSION")
	w.Header().Set("VERSION", fmt.Sprintf("%s", VERSION))
	for k, v := range r.Header {
		w.Header().Set(k, fmt.Sprintf("%s", v))
	}
	io.WriteString(w, "客户端地址：端口"+r.RemoteAddr+"\n")
	io.WriteString(w, "http返回码:200")
}
