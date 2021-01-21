package client

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"strings"

	"github.com/MouseHatGames/hat/pkg/client/proto"
	"google.golang.org/grpc"
)

type Client interface {
	io.Closer

	// Get fetches a value from the specified path.
	Get(path ...string) Value

	// GetBulk fetches multiple values from the store simultaneously.
	GetBulk(paths [][]string) ([]Value, error)

	// Set converts a value into a JSON string and stores it.
	Set(val interface{}, path ...string) error

	// SetRaw stores a JSON-encoded value.
	SetRaw(val string, path ...string) error

	// Del deletes a value.
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

// SplitPath splits a path (i.e. foo.bar.test) into its parts (i.e. ["foo", "bar", "test"])
func SplitPath(path string) []string {
	return strings.Split(path, ".")
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

func (c *client) GetBulk(paths [][]string) ([]Value, error) {
	req := proto.BulkRequest{
		Paths: make([]*proto.Path, len(paths)),
	}

	for i, p := range paths {
		req.Paths[i] = &proto.Path{Parts: p}
	}

	resp, err := c.cl.GetBulk(context.Background(), &req)
	if err != nil {
		return nil, err
	}

	values := make([]Value, len(resp.Data))
	for i, v := range resp.Data {
		values[i] = &jsonValue{val: v.Json}
	}

	return values, nil
}

func (c *client) Set(val interface{}, path ...string) error {
	str, err := json.Marshal(val)
	if err != nil {
		return fmt.Errorf("json encode: %w", err)
	}

	return c.SetRaw(string(str), path...)
}

func (c *client) SetRaw(val string, path ...string) error {
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
