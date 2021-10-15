package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	_ "net/http/pprof"
	"os"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", rootHandler)
	mux.HandleFunc("/healthz", healthz)
	err := http.ListenAndServe(":8888", mux)
	if err != nil {
		log.Fatal(err)
	}
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
