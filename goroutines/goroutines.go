package main

import (
	"fmt"
	"strconv"
	"sync"
)

func runMe(name string) {
	fmt.Println("Hello from goroutine:", name)
}

func main() {
	var wg sync.WaitGroup
	// wg.Add(1)
	// go func(name string) {
	// 	runMe(name)
	// 	wg.Done()
	// }("Vetri")
	var t string
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			t = strconv.Itoa(i)
			runMe(t)
			wg.Done()
		}(i)
	}

	wg.Wait()
	// time.Sleep(1 * time.Second)

}
