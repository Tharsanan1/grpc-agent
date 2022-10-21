generate_go:
	mkdir -p temp
	protoc --go-grpc_out=temp --go_out=temp --go_opt=paths=source_relative --go-grpc_opt=paths=source_relative protos/swagger.proto
	mv temp/protos/swagger*.go grpc/swagger/
	protoc --go-grpc_out=temp --go_out=temp --go_opt=paths=source_relative --go-grpc_opt=paths=source_relative protos/common.proto
	mv temp/protos/common*.go grpc/common/
	protoc --go-grpc_out=temp --go_out=temp --go_opt=paths=source_relative --go-grpc_opt=paths=source_relative protos/api.proto
	mv temp/protos/api*.go grpc/api/
	rm -rf temp/