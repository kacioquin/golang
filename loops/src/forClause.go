package forclause

import "fmt"

// func forclause() {
// 	forClauseIncrementByN(0, 10, 2)
// }

func ForClauseIncrementBy1(initial, condition int) {
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
