FROM docker.uclv.cu/golang:1.20-alpine

WORKDIR /app/rpc

RUN go install github.com/cosmtrek/air@latest

RUN go clean -modcache

COPY ./src/rpc/go.mod ./

RUN go mod download

RUN go clean -modcache