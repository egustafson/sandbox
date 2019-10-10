.PHONY: clean all deps build run

all:
	@echo ""
	@echo "make (deps | build | run | clean)"
	@echo ""

deps:
	glide init
	glide update
	glide install

build: consumer producer

consumer:
	go build consumer.go

producer:
	go build producer.go

clean:
	rm -rf vendor glide.yaml glide.lock producer consumer
