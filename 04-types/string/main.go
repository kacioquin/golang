package main

import "fmt"

func main() {
	s := "Hello world in GoLang"

	fmt.Println(s)
	fmt.Printf("%T", s)
	fmt.Println("")

	bs := []byte(s)
	fmt.Println(bs)
	fmt.Printf("%T\n", bs)

	for i := 0; i < len(s); i++ {
		fmt.Printf("%#U", s[i])
	}

}
