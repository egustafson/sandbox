FROM golang:1.8

## Place the target application in a directory named 'app'
WORKDIR /go/src/app
COPY . .

RUN go get -d -v ./...
RUN go install -v ./...

## "install" the (implicitly built) executable, 'app'.
## The exec is named 'app' because the parent dir is 'app'

## Run the executable named 'app' (w/ no flags) by default.
CMD ["app"]
