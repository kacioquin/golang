package main

import (
	"fmt"
	"log"

	"github.com/kacioquin/golang/tutorial/greetings"
)

func main() {
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	names := []string{"John", "Paul", "Grag"}

	messages, err := greetings.Hellos(names)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(messages)
}
