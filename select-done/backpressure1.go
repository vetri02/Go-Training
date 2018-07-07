package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	workers := 20000
	pool := make(chan func(int) int, workers)
	for i := 0; i < workers; i++ {
		pool <- func(in int) int {
			time.Sleep(1 * time.Second)
			result := 2 * in
			return result
		}
	}

	var wg sync.WaitGroup
	count := 0
	totalStart := time.Now()
	for i := 0; i < 100000; i++ {
		start := time.Now()
		select {
		case f := <-pool:
			fmt.Println("processing", i)
			count++
			wg.Add(1)
			go func(in int) {
				out := f(in)
				fmt.Println("got", out, "for", in, "after", time.Now().Sub(start))
				wg.Done()
			}(i)
		default:
			fmt.Println("rejected request", i, "too busy")
		}

	}
	wg.Wait()
	fmt.Println("totalProcessed", count)
	fmt.Println("totaltime", time.Now().Sub(totalStart))
}
