package randomwalk

import "time"

type Option func(*randomWalkOptions)

func WithRate(d time.Duration, deviation time.Duration) Option {
	return func(rw *randomWalkOptions) {
		rw.rate = d
		rw.rateDeviation = deviation
	}
}

func OnlyPositive() Option {
	return func(rw *randomWalkOptions) {
		rw.polarity = Positive
	}
}

func OnlyNegative() Option {
	return func(rw *randomWalkOptions) {
		rw.polarity = Negative
	}
}

func Min(x float64) Option {
	return func(rw *randomWalkOptions) {
		rw.min = x
	}
}

func Max(x float64) Option {
	return func(rw *randomWalkOptions) {
		rw.min = x
	}
}
