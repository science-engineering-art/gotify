module github.com/science-engineering-art/gotify/api

go 1.20

require (
	github.com/gofiber/fiber/v2 v2.46.0
	github.com/google/uuid v1.3.0
	go.mongodb.org/mongo-driver v1.11.7
	google.golang.org/protobuf v1.30.0
)

require (
	github.com/cespare/xxhash/v2 v2.2.0 // indirect
	github.com/dgryski/go-rendezvous v0.0.0-20200823014737-9f7001d12a5f // indirect
	github.com/dhowden/tag v0.0.0-20220618230019-adf36e896086 // indirect
	github.com/golang/protobuf v1.5.3 // indirect
	github.com/golang/snappy v0.0.4 // indirect
	github.com/kr/text v0.2.0 // indirect
	github.com/montanaflynn/stats v0.7.1 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/redis/go-redis/v9 v9.0.5 // indirect
	github.com/science-engineering-art/kademlia-grpc v0.4.7 // indirect
	github.com/xdg-go/pbkdf2 v1.0.0 // indirect
	github.com/xdg-go/scram v1.1.2 // indirect
	github.com/xdg-go/stringprep v1.0.4 // indirect
	github.com/youmark/pkcs8 v0.0.0-20201027041543-1326539a0a0a // indirect
	golang.org/x/crypto v0.9.0 // indirect
	golang.org/x/net v0.10.0 // indirect
	golang.org/x/sync v0.2.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20230530153820-e85fd2cbaebc // indirect
)

require (
	github.com/andybalholm/brotli v1.0.5 // indirect
	github.com/klauspost/compress v1.16.5 // indirect
	github.com/mattn/go-colorable v0.1.13 // indirect
	github.com/mattn/go-isatty v0.0.19 // indirect
	github.com/mattn/go-runewidth v0.0.14 // indirect
	github.com/philhofer/fwd v1.1.2 // indirect
	github.com/rivo/uniseg v0.4.4 // indirect
	github.com/savsgio/dictpool v0.0.0-20221023140959-7bf2e61cea94 // indirect
	github.com/savsgio/gotils v0.0.0-20230208104028-c358bd845dee // indirect
	github.com/tinylib/msgp v1.1.8 // indirect
	github.com/valyala/bytebufferpool v1.0.0 // indirect
	github.com/valyala/fasthttp v1.47.0 // indirect
	github.com/valyala/tcplisten v1.0.0 // indirect
	golang.org/x/sys v0.8.0 // indirect
	golang.org/x/text v0.9.0 // indirect
	google.golang.org/grpc v1.55.0
)

require (
	github.com/gofrs/uuid v4.4.0+incompatible
	github.com/science-engineering-art/gotify/peer v0.0.0-00010101000000-000000000000
	github.com/science-engineering-art/gotify/tracker v0.0.0-00010101000000-000000000000
)

replace github.com/science-engineering-art/gotify/peer => ../peer

replace github.com/science-engineering-art/gotify/tracker => ../tracker

replace github.com/science-engineering-art/kademlia-grpc => ../../kademlia-grpc
