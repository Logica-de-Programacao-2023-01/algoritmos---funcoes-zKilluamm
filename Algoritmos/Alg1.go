package main

import (
	"fmt"
)

func calcularMedia(slice []int) float64 {
	soma := 0
	for _, valor := range slice {
		soma += valor
	}

	media := float64(soma) / float64(len(slice))
	return media
}

func main() {
	var tamanho int
	fmt.Print("Digite o tamanho do slice: ")
	fmt.Scanln(&tamanho)

	slice := make([]int, tamanho)
	for i := 0; i < tamanho; i++ {
		fmt.Printf("Digite o valor %d: ", i+1)
		fmt.Scanln(&slice[i])
	}

	media := calcularMedia(slice)
	fmt.Println("A média é:", media)
}
