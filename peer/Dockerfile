FROM docker.uclv.cu/golang:1.20-alpine

WORKDIR /peer

COPY ./ ./

# EXPOSE 8080
# EXPOSE 32140/udp

CMD [ "go", "run", "main.go" ]

# RUN go install github.com/cosmtrek/air@v1.43.0

# RUN go clean -modcache

# COPY ./src/peer/go.mod ./

# RUN go mod download

# CMD ["air", "-c", ".air.toml"]