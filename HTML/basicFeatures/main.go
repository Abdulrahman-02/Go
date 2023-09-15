package main

import (
	"fmt"
	"math/rand"
)

func main() {
	const Watersports = iota
	fmt.Println("Watersports: ", Watersports)
	fmt.Println("My value: ",rand.Int())

	var a int = 10
	var b *int = &a

	fmt.Println("a: ", a)
	fmt.Println("a address: ", &a)
	fmt.Println("b: ", b)
	fmt.Println("b address: ", &b)
	fmt.Println("b value: ", *b)

	// page 87



}
