module github.com/science-engineering-art/gotify/tracker

go 1.20

require (
	github.com/jbenet/go-base58 v0.0.0-20150317085156-6237cf65f3a6
	google.golang.org/grpc v1.55.0
)

require github.com/science-engineering-art/kademlia-grpc v0.0.0-00010101000000-000000000000

require (
	github.com/golang/protobuf v1.5.3 // indirect
	golang.org/x/net v0.10.0 // indirect
	golang.org/x/sys v0.8.0 // indirect
	golang.org/x/text v0.9.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20230530153820-e85fd2cbaebc // indirect
	google.golang.org/protobuf v1.30.0 // indirect
)

replace github.com/science-engineering-art/kademlia-grpc => ../../kademlia-grpc
