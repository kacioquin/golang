package main

import "fmt"

//Go is static language, with static typing...
//strings
var z string = "I'm a string type variable"
var x = "I'm a string type variable too, but declared of different way "

//int
var y = 10 //I'm a int type variable

func main() {
	fmt.Println("Present variables")

	fmt.Printf("variable x - %T\n", x)

	fmt.Printf("variable y - %T\n", y)

	fmt.Printf("variable z - %T\n", z)

}
