package randomwalk_test

import (
	"context"
	"fmt"
	"time"

	"github.com/draeder/randomwalk-go"
)

func ExampleRandomWalk() {
	// Provides a context to cancel the randomwalk stream after 5 seconds.
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// A set of optional arguments that may be passed into randomwalk.
	opts := []randomwalk.RandomWalkOption{
		randomwalk.WithRate(30*time.Millisecond, 10*time.Millisecond),
		randomwalk.OnlyPositive(),
		randomwalk.Min(0.00000000001),
		randomwalk.Max(1000),
	}

	walk := randomwalk.RandomWalk(ctx, 40.0, 5.10, opts...)

	for n := range walk {
		fmt.Println(n)
	}
}
