package randomwalk

import (
	"context"
	"math"
	"math/rand"
	"time"
)

type BoxMüller struct {
	r     *rand.Rand
	mu    float64
	sigma float64
}

type randomWalkOptions struct {
	rate          time.Duration
	rateDeviation time.Duration
	polarity      Polarity
	min           float64
	max           float64
}

type Polarity int

const (
	Both Polarity = iota
	Positive
	Negative
)

func New(mu, sigma float64) *BoxMüller {
	return &BoxMüller{
		r:     rand.New(rand.NewSource(time.Now().UnixNano())),
		mu:    mu,
		sigma: sigma,
	}
}

func (bm *BoxMüller) Rand() (float64, float64) {
	x := bm.r.Float64()
	y := bm.r.Float64()
	z1 := math.Sqrt(-2.0*math.Log(x)) * math.Cos(2.0*math.Pi*y)
	z2 := math.Sqrt(-2.0*math.Log(x)) * math.Sin(2.0*math.Pi*y)
	return bm.sigma*z1 + bm.mu, bm.sigma*z2 + bm.mu
}

func (bm *BoxMüller) RandLimits(p Polarity, min float64, max float64) (float64, float64) {

	z1, z2 := bm.Rand()

	switch p {
	case Positive:
		bm.mu = math.Max(z2, min)
		return math.Max(z1, min), bm.mu
	case Negative:
		bm.mu = math.Min(z2, -max)
		return math.Min(z1, -max), bm.mu
	default:
		return z1, z2
	}
}

//func RandomWalk(ctx context.Context, mean float64, sd float64, rate time.Duration, rateDeviation time.Duration, polarity Polarity) <-chan float64 {
func RandomWalk(ctx context.Context, mean float64, sd float64, opts ...Option) <-chan float64 {

	options := &randomWalkOptions{}

	for _, opt := range opts {
		opt(options)
	}

	min := int64(options.rate - options.rateDeviation)
	max := int64(options.rate + options.rateDeviation)

	rw := make(chan float64, 0)

	go func() {
		defer close(rw)

		bm := New(mean, sd)

		for {
			select {
			case <-time.After(time.Duration(rand.Int63n(max-min+1) + min)):
			case <-ctx.Done():
				return
			}

			z1, z2 := bm.RandLimits(options.polarity, options.min, options.max)
			bm.mu = z2

			rw <- z1

		}
	}()

	return rw

}
