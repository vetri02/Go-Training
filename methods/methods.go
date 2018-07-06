package main

import (
	"fmt"
)

type Foo struct {
	A int
	B string
}

func (g Foo) String() string {
	return fmt.Sprintf("A: %d and B: %s", g.A, g.B)
}

func (g *Foo) Double() {
	g.A = g.A * 2
}

func main() {
	f := Foo{
		A: 10,
		B: "Hello",
	}
	fmt.Println(f.String())
	f.Double()
	fmt.Println(f.String())
}
