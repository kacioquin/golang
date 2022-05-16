package main

import "fmt"

func main() {
	x := 42 //Declara a variável x e atribui o valor 42 a ela
	fmt.Println(x)
	x = 99 //Nesse ponto apenas atribui o valor 99 a variavel x, visto que ela ja havia sido declarada 2 passos acima
	// com o operador de declaracao curta -> :=
	fmt.Println(x)
	name := "teste"
	fmt.Println(name)
}
