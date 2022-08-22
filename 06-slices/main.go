package main

import "fmt"

func main() {

	var y [5]int
	fmt.Println(y)

	//SLICE - composite literal
	x := []int{1, 2, 3, 4, 5}

	forWithRange(x)
	forLoop(x)
	slicingASlice(x, 0)

}

func forWithRange(myVar []int) {
	fmt.Println("Loop using keyword range")
	for i, v := range myVar {
		fmt.Println("index", i, "value", v)
	}
}

func forLoop(myVar []int) {
	fmt.Println("Loop with tradicional For")
	for i := 0; i < len(myVar); i++ {
		fmt.Println("index", i, "value", myVar[i])
	}
}

func slicingASlice(myVar []int, initial int) {
	fmt.Println("Slicing a slice")
	fmt.Println("Original slice", myVar)
	fmt.Println("Parameters for slicing", initial)
	fmt.Println("New slice")
	newSlice := myVar[initial:]
	fmt.Println("Result", newSlice)
}
