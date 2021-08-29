package randomwalk

import (
	"encoding/binary"
	"io"
	"math/rand"
)

type randomSource struct {
	r io.Reader
}

// newRandomSource creates a randomSource reader
func newRandomSource(r io.Reader) rand.Source {
	return &randomSource{
		r: r,
	}
}

// Seed .
func (src randomSource) Seed(seed int64) {}

// Int63 returns a default seed if none is provided.
func (src randomSource) Int63() int64 {
	return int64(src.Uint64() & ^uint64(1<<63))
}

// Uint64 .
func (src randomSource) Uint64() (v uint64) {
	err := binary.Read(src.r, binary.BigEndian, &v)
	if err != nil {
		panic(err)
	}
	return v
}
