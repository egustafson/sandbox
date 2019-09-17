all:
	@echo "\n  make build | make run\n"

build:
	docker build -t webapp .

run:
	docker run -it --rm --name webapp -p 5000:5000 webapp

runenv:
	docker run -it --rm --name webapp -p 5000:5000 -e 'DOCKER_TURD=set-in-makefile' webapp
