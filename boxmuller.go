package randomwalk

import (
	"math"
	"math/rand"
	"time"
)

// Polarity is type that describes whether RandomWalk returns positive, negative or both types of random numbers.
type Polarity int

const (
	Both Polarity = iota
	Positive
	Negative
)

// BoxMüller holds the state of the BoxMüller transform.
type BoxMüller struct {
	rand  *rand.Rand
	mu    float64
	sigma float64
}

// New initializes and returns a Box-Müller transform.
func New(mu, sigma float64, opts ...BoxMüllerOption) *BoxMüller {
	bm := &BoxMüller{
		mu:    mu,
		sigma: sigma,
	}

	for _, opt := range opts {
		opt(bm)
	}

	if bm.rand == nil {
		bm.rand = rand.New(rand.NewSource(time.Now().UnixNano()))
	}

	return bm

}

// SetMu sets the mu (mean) of the BoxMuller transform
func (bm *BoxMüller) SetMu(x float64) {
	bm.mu = x
}

// Rand returns Box-Müller number pairs.
func (bm *BoxMüller) Rand() (float64, float64) {
	x := bm.rand.Float64()
	y := bm.rand.Float64()
	z1 := math.Sqrt(-2.0*math.Log(x)) * math.Cos(2.0*math.Pi*y)
	z2 := math.Sqrt(-2.0*math.Log(x)) * math.Sin(2.0*math.Pi*y)
	return bm.sigma*z1 + bm.mu, bm.sigma*z2 + bm.mu
}

// RandLimits returns Box-Müller number pairs that are within the specified limits.
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
