export tag=v1.6
root:
	export ROOT=github.com/yagamishi/golang

build:
	echo "building httpserver binary"
	mkdir -p bin/amd64
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o bin/amd64/httpserver .

release: build
	echo "building httpserver container"
	docker build -t yagamishi/httpserver:${tag} .

push: release
	echo "pushing yagamishi/httpserver"
	docker push yagamishi/httpserver:${tag}
