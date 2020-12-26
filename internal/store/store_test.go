package store

import (
	"testing"

	"github.com/MouseHatGames/hat/internal/proto"
	"github.com/stretchr/testify/assert"
)

func TestJoinPath(t *testing.T) {
	t.Run("one part", func(t *testing.T) {
		path := &proto.Path{Parts: []string{"one"}}

		joined := string(joinPath(path))

		assert.Equal(t, "one", joined)
	})

	t.Run("three parts", func(t *testing.T) {
		path := &proto.Path{Parts: []string{"one", "two", "three"}}

		joined := string(joinPath(path))

		assert.Equal(t, "one/two/three", joined)
	})
}
