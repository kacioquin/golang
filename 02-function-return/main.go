package main

import "fmt"

func main() {
	n, err := fmt.Println("Hello, I am a function with return")
	fmt.Println("I'm a function return - ", n)
	fmt.Println("I'm a error list - ", err)
}
