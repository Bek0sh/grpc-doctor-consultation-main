db:
	docker start postgres15

proto:
	protoc --go_out=. --go-grpc_out=.  ./api/proto/userinfo.proto 