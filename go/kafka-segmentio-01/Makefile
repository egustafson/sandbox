.PHONY: clean all deps build run

all:
	@echo ""
	@echo "make (deps | build | run | clean)"
	@echo ""

deps: glide.yaml
	glide update
	glide install

glide.yaml:
	glide init

build: reader writer

reader:
	go build reader.go

writer:
	go build writer.go

clean:
	rm -rf vendor glide.yaml glide.lock reader writer
