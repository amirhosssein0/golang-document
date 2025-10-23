package main

import "fmt"

func main() {
	fmt.Print("hi", "amirhossein")
	fmt.Print("aaa")
	fmt.Println("hi", "amirhossein")
	fmt.Println("aaa")
	fmt.Printf("hi %s test", "amirhossein")
	// fmt.Printf("hi %s %d %T %f %v test",)

	var a string = "test"
	b := "test2"

	fmt.Println(a == b)

	c := a + " " + b

	fmt.Println(c)

}