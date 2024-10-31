# Makefile

TAG = 0.1.0

hello: main.go
	go build -o hello ./...

build: hello

package: hello Dockerfile
	docker build -t hello-docker:$(TAG) -f Dockerfile .
