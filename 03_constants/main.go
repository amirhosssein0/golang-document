package main

import (
	"fmt"
	//"math"
)

func main() {
	// Declaring a constant
	const a = "test1"
	fmt.Println(a)

	// Declaring a group of constants
	const (
		c = 1
		d = 2
		e = 3
	)

	// The value of a constant should be known at compile time
	//We can't value a constant by a function
	// var a = math.Sqrt(4)
	// fmt.Println(a)
	// const b = math.Sqrt(4)
	// fmt.Println(b)

	// Typed and Untyped Constants
	const name bool = true
	var f = name
	fmt.Println(name, f)
}