# -*- Dockerfile -*-
# -------------------------------------
# Build Stage
#
FROM golang:latest AS builder

WORKDIR /workspace

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .
RUN make build

# -------------------------------------
# Package Stage
#
FROM debian:trixie-slim

WORKDIR /app
COPY --from=builder /workspace/dist/kubever /app/kubever
EXPOSE 8080/tcp
CMD [ "/app/kubever" ]
