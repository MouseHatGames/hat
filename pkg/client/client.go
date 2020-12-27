package client

import (
	"context"
	"fmt"
	"io"

	"github.com/MouseHatGames/hat/internal/proto"
	"google.golang.org/grpc"
)

type Client interface {
	io.Closer

	Get(path ...string) Value
	Set(val string, path ...string) error
	Del(path ...string) error
}

func Dial(addr string) (Client, error) {
	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		return nil, fmt.Errorf("dial: %w", err)
	}

	return &client{
		conn: conn,
		cl:   proto.NewHatClient(conn),
	}, nil
}

type client struct {
	conn *grpc.ClientConn
	cl   proto.HatClient
}

func (c *client) Close() error {
	return c.conn.Close()
}

func (c *client) Get(path ...string) Value {
	d, err := c.cl.Get(context.Background(), &proto.Path{Parts: path})
	if err != nil {
		return &jsonValue{err: err}
	}
	return &jsonValue{val: d.Json}
}

func (c *client) Set(val string, path ...string) error {
	_, err := c.cl.Set(context.Background(), &proto.SetRequest{
		Path:  &proto.Path{Parts: path},
		Value: &proto.Data{Json: val},
	})
	return err
}

func (c *client) Del(path ...string) error {
	_, err := c.cl.Delete(context.Background(), &proto.Path{Parts: path})
	return err
}
