package client

import (
	"testing"
)

func BenchmarkGet(b *testing.B) {
	const path = "test"

	cl, err := Dial("localhost:4659")
	if err != nil {
		b.Errorf("failed to dial server: %s", err)
		return
	}
	cl.Set("123", path)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		cl.Get(path)
	}
}
