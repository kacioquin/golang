package main

import (
	"fmt"
)

type hotdog int

var x hotdog
var y int

func main() {
	fmt.Printf("%v\n", x)
	fmt.Printf("%T\n", x)
	x = 42
	fmt.Printf("%v\n", x)

	y = int(x)

	fmt.Println(y)
	fmt.Printf("%T", y)
}
