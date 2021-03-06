# randomwalk-go
Creates a trend-oriented stream of random numbers using a Box Müller transform. `randomwalk-go` is an implementation of [random-walk](https://www.npmjs.com/package/random-walk) in go.


## Install
```
go get github.com/draeder/randomwalk-go
```

## Example
```go
package main

import (
	"context"
	"fmt"
	"time"

	"github.com/draeder/randomwalk-go"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	opts := []randomwalk.Option{
		randomwalk.WithRate(30*time.Millisecond, 10*time.Millisecond),
		randomwalk.OnlyPositive(),
		randomwalk.Min(0.00000000001),
		randomwalk.Max(1000),
	}

	walk := randomwalk.RandomWalk(ctx, 40.0, 5.10, opts...)

	for n := range walk {
		fmt.Println(n)
	}

	fmt.Println(ctx.Err())
}
```
