FROM docker.uclv.cu/golang:1.20-alpine

WORKDIR /app

RUN go install github.com/cosmtrek/air@latest

RUN mkdir rpc
COPY ./src/rpc/go.mod ./rpc/

RUN mkdir api
COPY ./src/api/go.mod ./api/

COPY ./src/docker/api/go.work ./

RUN cd rpc
RUN go mod download
RUN cd ..

RUN cd api
RUN go mod download

WORKDIR /app/api

CMD ["air", "-c", ".air.toml"]