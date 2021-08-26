package randomwalk

import (
	"math"
	"math/rand"
	"time"
)

type BoxMüller struct {
	r     *rand.Rand
	mu    float64
	sigma float64
}

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

func RandomWalk(mean float64, sd float64, rate time.Duration, rateDeviation time.Duration) <-chan float64 {

	min := int64(rate - rateDeviation)
	max := int64(rate + rateDeviation)

	rw := make(chan float64, 0)

	go func() {
		defer close(rw)
		for {
			<-time.After(time.Duration(rand.Int63n(max-min+1) + min))
			bm := New(mean, sd)
			z1, z2 := bm.Rand()
			_ = z1
			mean = z2

			rw <- z1

			//fmt.Println("\033[H\033[2J")
			//fmt.Println(math.Round(z1*100) / 100)

		}
	}()

	return rw

}
