package main

import (
	"fmt"
)

// Foo contains int and string
type Foo struct {
	A int
	b string
}

func main() {
	f := Foo{}
	fmt.Println(f)

	g := Foo{10, "hello"}
	fmt.Println(g)

	h := Foo{
		b: "go",
	}
	fmt.Println(h)

	h.A = 12
	fmt.Println(h)
	fmt.Println(h.A)
	fmt.Println(h.b)

	i := Foo{
		A: 2,
	}

	//i1 := i
	var i1 Foo
	i1 = i

	fmt.Println(i1)
	fmt.Println(i)

	var f3 *Foo = &g
	f3.A = 234
	fmt.Println(f3)
	fmt.Println(f3.A)
	fmt.Println(g.A)

}
