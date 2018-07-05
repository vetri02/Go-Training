package main

import (
	"fmt"
)

type Foo struct {
	A int
	B string
}

type Bar struct {
	C string
	F Foo
}

type Baz struct {
	D string
	Foo
}

func main() {
	f := Foo{A: 10, B: "Hello"}
	fmt.Println(f)

	b := Bar{
		C: "good",
		F: f,
	}

	fmt.Println(b)

	c := Baz{
		D:   "star",
		Foo: f,
	}

	fmt.Println(c)

	var f2 Foo = c.Foo
	fmt.Println(f2)
}
