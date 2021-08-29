package randomwalk

import (
	"encoding/binary"
	"io"
	"math/rand"
)

type randomSource struct {
	r io.Reader
}

func newRandomSource(r io.Reader) rand.Source {
	return &randomSource{
		r: r,
	}
}

func (src randomSource) Seed(seed int64) {}

func (src randomSource) Int63() int64 {
	return int64(src.Uint64() & ^uint64(1<<63))
}

func (src randomSource) Uint64() (v uint64) {
	err := binary.Read(src.r, binary.BigEndian, &v)
	if err != nil {
		panic(err)
	}
	return v
}
