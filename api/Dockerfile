FROM docker.uclv.cu/golang:1.20-alpine

WORKDIR /api

COPY ./ ./

EXPOSE 80
# EXPOSE 53123/udp

# EXPOSE 8080
# EXPOSE 32140/udp

CMD ["go", "run", "main.go"]

# RUN go install github.com/cosmtrek/air@v1.43.0

# RUN go clean -modcache

# COPY ./src/api/go.mod ./

# RUN go mod download

# CMD ["air", "-c", ".air.toml"]
