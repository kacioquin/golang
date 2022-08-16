package main

import "fmt"

func ForClauseIncrementBy1(initial int, condition int) {
	for i := initial; i < condition; i++ {
		fmt.Println("The value is: ", i)
	}
}

func ForClauseIncrementByN(initial, condition, increment int) {
	fmt.Printf("Increment by %v\n", increment)
	for i := initial; i <= condition; i += increment {
		fmt.Println("The value is: ", i)
	}
}
