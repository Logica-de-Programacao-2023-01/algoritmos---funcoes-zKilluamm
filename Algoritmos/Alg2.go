package main

import (
	"fmt"
)

func contarVogais(str string) int {
	vogais := []rune{'a', 'e', 'i', 'o', 'u'}
	contador := 0

	for _, char := range str {
		for _, vogal := range vogais {
			if char == vogal || char == vogal-32 {
				contador++
				break
			}
		}
	}

	return contador
}

func main() {
	var texto string
	fmt.Print("Digite uma frase: ")
	fmt.Scanln(&texto)

	quantidade := contarVogais(texto)
	fmt.Println("A quantidade de vogais é:", quantidade)
}
