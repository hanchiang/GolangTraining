package main

import "fmt"

func main() {
	// Shorthands, type inferred
	a := 10
	b := "Hello"
	c := 5.51
	d := true

	var e int64
	var f int64 = 12
	var g string = "World"
	var h string
	var i float64 = 12.4123
	var j float64
	var k bool = false
	var l bool

	fmt.Printf("%d: %T\n", a, a)
	fmt.Printf("%s: %T\n", b, b)
	fmt.Printf("%f: %T\n", c, c)
	fmt.Printf("%t: %T\n", d, d)

	fmt.Printf("%d: %T\n", e, e)
	fmt.Printf("%d: %T\n", f, f)
	fmt.Printf("%s: %T\n", g, g)
	fmt.Printf("%s: %T\n", h, h)
	fmt.Printf("%f: %T\n", i, i)
	fmt.Printf("%f: %T\n", j, j)
	fmt.Printf("%t: %T\n", k, k)
	fmt.Printf("%t: %T\n", l, l)
}
