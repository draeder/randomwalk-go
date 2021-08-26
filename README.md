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
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/draeder/randomwalk-go"
)

func main() {

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer cancel()

	opts := []randomwalk.Option{
		randomwalk.WithRate(30*time.Millisecond, 10*time.Millisecond),
		randomwalk.OnlyPositive(),
		randomwalk.Min(0.00000000001),
	}

	rw := randomwalk.RandomWalk(ctx, 40.0, 0.10, opts...)
	for n := range rw {

		fmt.Println("\033[H\033[2J")
		fmt.Println(n)
	}

	fmt.Println(ctx.Err())

}
```
