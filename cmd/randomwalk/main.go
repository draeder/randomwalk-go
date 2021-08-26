package main

import (
	"fmt"
	"time"

	"github.com/draeder/randomwalk-go"
)

func main() {
	rw := randomwalk.RandomWalk(1.0, 1.0, 300*time.Millisecond, 100*time.Millisecond)
	for n := range rw {
		if n <= 0.0 {
			continue
		}
		fmt.Println("\033[H\033[2J")
		fmt.Println(n)
	}
}
