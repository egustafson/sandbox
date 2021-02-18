.PHONEY: run unit-test integration-test

unit-test:
	go test -v ./...

integration-test:
	go test -v --tags=integration ./...

run:
	go run main.go
