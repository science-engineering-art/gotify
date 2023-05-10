FROM docker.uclv.cu/golang:1.20-alpine

WORKDIR /app

RUN go install github.com/cosmtrek/air@latest

COPY ./src/peer/go.mod ./
RUN go mod download

CMD ["air", "-c", ".air.toml"]