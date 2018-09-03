package main

import "fmt"

func main() {
	fmt.Println("Hello world!")

	// # = add leading 0 for oct and leading 0x for hex
	fmt.Printf("%d - %b - %#x - %q \n", 42, 42, 42, 42)
}
