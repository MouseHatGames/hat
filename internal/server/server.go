package server

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/MouseHatGames/hat/internal/config"
	"github.com/MouseHatGames/hat/internal/proto"
	"github.com/MouseHatGames/hat/internal/store"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func Start(cfg *config.Config, store store.Store) error {
	addr := fmt.Sprintf(":%d", cfg.Port)
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		return fmt.Errorf("listen: %w", err)
	}
	defer lis.Close()

	log.Printf("listening on %s", addr)

	grpcServer := grpc.NewServer()
	proto.RegisterHatServer(grpcServer, &hatServer{store: store})

	if err := grpcServer.Serve(lis); err != nil {
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

func (s *hatServer) Delete(ctx context.Context, req *proto.Path) (*proto.Empty, error) {
	if err := s.store.Del(req); err != nil {
		return nil, err
	}
	return empty, nil
}
