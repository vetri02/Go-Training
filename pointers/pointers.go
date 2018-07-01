package main

import (
	"fmt"
)

// * ---> value
// & ----> location

func setTo10(e *int) {
	*e = 10
}

func main() {
	a := 10
	b := &a
	c := a

	fmt.Println(a, b, *b, c)

	a = 20
	fmt.Println(a, b, *b, c)

	*b = 30
	fmt.Println(a, b, *b, c)

	c = 40
	fmt.Println(a, b, *b, c)

	///////////////////////

	d := new(int)

	fmt.Println(d)
	fmt.Println(*d)

	///////////////////////

	e := 15
	fmt.Println(e)
	setTo10(&e)
	fmt.Println(e)

}
