package main

import "fmt"

func EncontrarPosicao(slice []int, valor int) int {
	for i, num := range slice {
		if num == valor {
			return i
		}
	}
	return -1
}

func main() {
	slice := []int{10, 20, 30, 40, 50}
	valor := 30
	posicao := EncontrarPosicao(slice, valor)
	fmt.Printf("Valor %d encontrado na posição %d\n", valor, posicao)

	valor = 60
	posicao = EncontrarPosicao(slice, valor)
	fmt.Printf("Valor %d encontrado na posição %d\n", valor, posicao)
}
