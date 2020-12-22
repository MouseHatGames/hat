package store

import (
	"bytes"
	"errors"
	"io"
)

var ErrKeyNotFound = errors.New("key was not found")

type Path interface {
	GetParts() []string
}

type Store interface {
	io.Closer

	Get(p Path) (string, error)
	Set(p Path, v string) error
	Del(p Path) error
}

func joinPath(path Path) []byte {
	parts := path.GetParts()

	if len(parts) == 0 {
		return []byte{}
	}

	buf := bytes.NewBufferString(parts[0])

	for _, p := range parts[1:] {
		buf.WriteRune('/')
		buf.WriteString(p)
	}

	return buf.Bytes()
}
