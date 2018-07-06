package main

import (
	"fmt"
)

type myint int

func (mi myint) isEven() bool {
	return mi%2 == 0
}

func (mi *myint) Double() {
	*mi = *mi * 2
}

func main() {
	m := myint(10)
	fmt.Println(m)
	fmt.Println(m.isEven())
	m.Double()
	fmt.Println(m)

}
