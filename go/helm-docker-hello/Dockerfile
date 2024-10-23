FROM golang:bookworm as build-env

WORKDIR /go/src/app
ADD . /go/src/app/
RUN make build

FROM debian:bookworm-slim

COPY --from=build-env /go/src/app/hello /usr/local/bin
RUN apt update; apt upgrade -y; apt install -y vim curl

EXPOSE 8080/tcp
CMD ["hello"]
