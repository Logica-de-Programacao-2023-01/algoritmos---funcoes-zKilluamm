package main

import (
	"fmt"
	"strings"
)

func concatenarStrings(slice []string) string {
	resultado := strings.Join(slice, "")
	return resultado
}

func main() {
	var tamanho int
	fmt.Print("Digite o tamanho do slice: ")
	fmt.Scanln(&tamanho)

	slice := make([]string, tamanho)
	for i := 0; i < tamanho; i++ {
		fmt.Printf("Digite a string %d: ", i+1)
		fmt.Scanln(&slice[i])
	}

	concatenacao := concatenarStrings(slice)
	fmt.Println("A concatenação é:", concatenacao)
}
