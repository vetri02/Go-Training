package main

import (
	"fmt"
)

func multiples(i int) (chan int, chan struct{}) {
	out := make(chan int)
	done := make(chan struct{})
	curVal := 0
	go func() {
		for {
			select {
			case out <- curVal * i:
				curVal++
			case <-done:
				fmt.Println("goroutines are done")
				return
			}

		}
	}()
	return out, done
}

func main() {
	twosCh, done := multiples(2)
	for v := range twosCh {
		if v > 20 {
			break
		}
		fmt.Println(v)
	}
	close(done)

}
