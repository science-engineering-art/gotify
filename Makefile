
vendor:
	cd api && go mod tidy && go mod vendor && cd ..
	cd dns && go mod tidy && go mod vendor && cd ..
	cd peer && go mod tidy && go mod vendor && cd .. 
	cd tracker && go mod tidy && go mod vendor && cd .. 

cert: 
	cd api/cert && bash gen.sh && cd ../..
	cd peer/cert && bash gen.sh && cd ../.. 
	cd tracker/cert && bash gen.sh && cd ../.. 	

protoc:
	cd proto && protoc --go_out=../api/pb --go_opt=paths=source_relative --go-grpc_out=../api/pb --go-grpc_opt=paths=source_relative *.proto && cd ..

clean:
	docker rmi $(docker images | grep "<none>" | awk '{print $3}')
