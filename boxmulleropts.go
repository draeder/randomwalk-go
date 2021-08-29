package randomwalk

import (
	"io"
	"math/rand"
)

type BoxMüllerOption func(*BoxMüller)
// WithRandomSource implements a random source for Box-Muller.
func WithRandomSource(r io.Reader) BoxMüllerOption {
	return func(bm *BoxMüller) {
		bm.rand = rand.New(newRandomSource(r))
	}
}
