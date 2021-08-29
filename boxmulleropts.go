package randomwalk

import (
	"io"
	"math/rand"
)

type BoxMüllerOption func(*BoxMüller)

func WithRandomSource(r io.Reader) BoxMüllerOption {
	return func(bm *BoxMüller) {
		bm.rand = rand.New(newRandomSource(r))
	}
}
