all: proto

proto: pkg/proto/hat.pb.go
	@mkdir -p pkg/client/proto
	@cp internal/proto/*.go pkg/client/proto

pkg/proto/%.pb.go: internal/proto/%.proto
	@protoc -Iinternal/proto --go_out=internal/proto/ --go_opt=paths=source_relative --go-grpc_out=internal/proto/ --go-grpc_opt=paths=source_relative $<