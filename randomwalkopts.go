package randomwalk

import (
	"time"
)

type RandomWalkOption func(*randomwalkOptions)

func WithRate(d time.Duration, deviation time.Duration) RandomWalkOption {
	return func(rw *randomwalkOptions) {
		rw.rate = d
		rw.rateDeviation = deviation
	}
}

func OnlyPositive() RandomWalkOption {
	return func(rw *randomwalkOptions) {
		rw.polarity = Positive
	}
}

func OnlyNegative() RandomWalkOption {
	return func(rw *randomwalkOptions) {
		rw.polarity = Negative
	}
}

func Min(x float64) RandomWalkOption {
	return func(rw *randomwalkOptions) {
		rw.min = x
	}
}

func Max(x float64) RandomWalkOption {
	return func(rw *randomwalkOptions) {
		rw.min = x
	}
}

func WithNumberTransformer(transformer NumberTransformer) RandomWalkOption {
	return func(rw *randomwalkOptions) {
		rw.transformer = transformer
	}
}
