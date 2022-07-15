package main

import "fmt"

var x int
var y float64

func main() {

	x = 42
	y = 45.2312

	fmt.Println(x)
	fmt.Println(y)
	fmt.Printf("%T\n", x)
	fmt.Printf("%T\n", y)
}
