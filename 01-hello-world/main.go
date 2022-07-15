package main

import (
	"fmt"

	"github.com/kacioquin/golang/01-hello-world/hello"
)

func main() {
	message := hello.Hello()
	fmt.Println(message)
}
