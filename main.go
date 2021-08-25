package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

func main() {
	rate := 10
	interval := time.Tick(time.Duration(rate) * time.Millisecond)
	mean := 100.0
	sd := 0.01

	for range interval {
		bm := New(mean, sd)
		z1, z2 := bm.Rand()
		_ = z1
		mean = z2

		fmt.Println("\033[H\033[2J")
		fmt.Println(math.Round(z1*100) / 100)
		//log.Println(z1, z2)

	}

}

type BoxMüller struct {
	r     *rand.Rand
	mu    float64
	sigma float64
}

func New(mu, sigma float64) *BoxMüller {
	p := new(BoxMüller)
	p.r = rand.New(rand.NewSource(time.Now().UnixNano()))
	p.mu = mu
	p.sigma = sigma
	return p
}

func (bm *BoxMüller) Rand() (float64, float64) {
	x := bm.r.Float64()
	y := bm.r.Float64()
	z1 := math.Sqrt(-2.0*math.Log(x)) * math.Cos(2.0*math.Pi*y)
	z2 := math.Sqrt(-2.0*math.Log(x)) * math.Sin(2.0*math.Pi*y)
	return bm.sigma*z1 + bm.mu, bm.sigma*z2 + bm.mu
}
