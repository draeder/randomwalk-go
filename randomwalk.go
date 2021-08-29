// randomwalk implements a trend-oriented stream of random numbers using a Box Müller transform.
package randomwalk

import (
	"context"
	"math/rand"
	"time"
)

type NumberTransformer interface {
	RandLimits(p Polarity, min float64, max float64) (float64, float64)
	SetMu(x float64)
}

type randomwalkOptions struct {
	transformer   NumberTransformer
	rate          time.Duration
	rateDeviation time.Duration
	polarity      Polarity
	min           float64
	max           float64
}

// RandomWalk generates the trend-oriented stream of random numbers using a Box Müller transform,
// which will be sent through the returned channel.
func RandomWalk(ctx context.Context, mean float64, sd float64, opts ...RwOption) <-chan float64 {

	options := &randomwalkOptions{}

	for _, opt := range opts {
		opt(options)
	}

	if options.transformer == nil {
		options.transformer = New(mean, sd)
	}

	rw := make(chan float64, 0)

	min := int64(options.rate - options.rateDeviation)
	max := int64(options.rate + options.rateDeviation)

	go func() {
		defer close(rw)

		for {
			select {
			case <-time.After(time.Duration(rand.Int63n(max-min+1) + min)):
				z1, z2 := options.transformer.RandLimits(options.polarity, options.min, options.max)
				options.transformer.SetMu(z2)

				rw <- z1

			case <-ctx.Done():
				return
			}
		}
	}()

	return rw

}
