all: proto

proto: pkg/proto/hat.pb.go

pkg/proto/%.pb.go: proto/%.proto
	@protoc -Iproto --go_out=internal/proto/ --go_opt=paths=source_relative --go-grpc_out=internal/proto/ --go-grpc_opt=paths=source_relative $<