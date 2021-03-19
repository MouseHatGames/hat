package server

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/MouseHatGames/hat/internal/proto"
	"github.com/MouseHatGames/hat/internal/store"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func Start(port int, store store.Store) error {
	addr := fmt.Sprintf(":%d", port)
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		return fmt.Errorf("listen: %w", err)
	}
	defer lis.Close()

	log.Printf("listening on %s", addr)

	sv := grpc.NewServer()
	proto.RegisterHatServer(sv, &hatServer{store: store})

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGHUP, syscall.SIGTERM, syscall.SIGQUIT)
	go func() {
		<-sig
		sv.GracefulStop()
	}()

	if err := sv.Serve(lis); err != nil {
		return fmt.Errorf("grpc serve: %w", err)
	}

	return nil
}

type hatServer struct {
	proto.UnimplementedHatServer
	store store.Store
}

var empty = &proto.Empty{}

func (s *hatServer) Set(ctx context.Context, req *proto.SetRequest) (*proto.Empty, error) {
	if err := s.store.Set(req.Path, req.Value.Json); err != nil {
		return nil, err
	}
	return empty, nil
}

func (s *hatServer) Get(ctx context.Context, req *proto.Path) (*proto.Data, error) {
	val, err := s.store.Get(req)
	if err != nil {
		if err == store.ErrKeyNotFound {
			return nil, status.Error(codes.NotFound, "key not found")
		}
		return nil, err
	}

	return &proto.Data{Json: val}, nil
}

func (s *hatServer) GetBulk(ctx context.Context, req *proto.BulkRequest) (*proto.BulkResponse, error) {
	paths := make([]store.Path, len(req.Paths))
	for i, v := range req.Paths {
		paths[i] = v
	}

	values, err := s.store.GetBulk(paths)
	if err != nil {
		return nil, err
	}

	resp := &proto.BulkResponse{
		Data: make([]*proto.Data, len(values)),
	}
	for i, v := range values {
		if v != nil {
			resp.Data[i] = &proto.Data{
				Json: *v,
			}
		}
	}

	return resp, nil
}

func (s *hatServer) Delete(ctx context.Context, req *proto.Path) (*proto.Empty, error) {
	if err := s.store.Del(req); err != nil {
		return nil, err
	}
	return empty, nil
}
