all: proto

proto: pkg/proto/hat.pb.go

pkg/proto/%.pb.go: proto/%.proto
	@protoc -Iproto --go_out=pkg/proto/ --go_opt=paths=source_relative --go-grpc_out=pkg/proto/ --go-grpc_opt=paths=source_relative $<