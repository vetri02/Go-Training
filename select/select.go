package main

import (
	"fmt"
)

func main() {
	in := make(chan int, 1)
	out := make(chan int, 1)

	out <- 1

	//DEADLOCK

	// in <- 2
	// fmt.Println("wrote 2 in 1")
	// v := <-out
	// fmt.Println("read", v, "from out")

	select {
	case in <- 2:
		fmt.Println("wrote 2 in 1")
	case v := <-out:
		fmt.Println("read", v, "from out")
	default:
		fmt.Println("nnnothinng")
	}

}
