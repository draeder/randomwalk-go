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

	//ctx2, cancel2 := context.WithDeadline(ctx, time.Now().Add(1*time.Minute))
	//defer cancel2()

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
