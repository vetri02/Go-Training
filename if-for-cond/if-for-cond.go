package main

import (
	"fmt"
)

func main() {

	//if cond

	// a := 10
	// if a > 5 {
	// 	fmt.Println("a is greater than 5")
	// } else {
	// 	fmt.Println("a is less than 5")
	// }

	// b := 10
	// if c := b / 2; b > c { // block scope
	// 	fmt.Println("b is greater than c:", c)
	// } else {
	// 	fmt.Println("b is less than c:", c)
	// }

	//for cond

	// for i := 0; i < 10; i++ {
	// 	// if i > 5 {
	// 	// 	break
	// 	// }

	// 	k := 3

	// 	if i == k {
	// 		continue
	// 	}
	// 	fmt.Println(i)
	// }

	// while

	// m := 0
	// for m < 10 {
	// 	fmt.Println(m)
	// 	m = m + 1
	// }

	// fmt.Println(m)

	// n := 0
	// for {
	// 	fmt.Println(n)
	// 	n = n + 1
	// 	if n > 10 {
	// 		break
	// 	}
	// }

	// fmt.Println(n)

	s := "hello world"
	for k, v := range s {
		fmt.Println(k, v, string(v))
	}
}
