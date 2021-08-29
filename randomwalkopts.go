package randomwalk

import (
	"time"
)

type RwOption func(*randomwalkOptions)

func WithRate(d time.Duration, deviation time.Duration) RwOption {
	return func(rw *randomwalkOptions) {
		rw.rate = d
		rw.rateDeviation = deviation
	}
}

func OnlyPositive() RwOption {
	return func(rw *randomwalkOptions) {
		rw.polarity = Positive
	}
}

func OnlyNegative() RwOption {
	return func(rw *randomwalkOptions) {
		rw.polarity = Negative
	}
}

func Min(x float64) RwOption {
	return func(rw *randomwalkOptions) {
		rw.min = x
	}
}

func Max(x float64) RwOption {
	return func(rw *randomwalkOptions) {
		rw.min = x
	}
}

func WithNumberTransformer(transformer NumberTransformer) RwOption {
	return func(rw *randomwalkOptions) {
		rw.transformer = transformer
	}
}
