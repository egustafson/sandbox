all:
	@echo "\n  make build | make run\n"

build:
	docker build -t go-webapp .

run:
	docker run -it --rm --name webapp -p 8080:8080 go-webapp
