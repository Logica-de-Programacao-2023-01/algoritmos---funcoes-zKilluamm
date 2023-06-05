package main

import (
	"errors"
	"fmt"
)

func IntersecaoSlices(slice1, slice2 []int) ([]int, error) {
	if len(slice1) == 0 || len(slice2) == 0 {
		return nil, errors.New("Um dos slices está vazio")
	}

	intersecao := make([]int, 0)
	ocorrencias := make(map[int]bool)

	for _, num := range slice1 {
		ocorrencias[num] = true
	}

	for _, num := range slice2 {
		if ocorrencias[num] {
			intersecao = append(intersecao, num)
		}
	}

	return intersecao, nil
}

func main() {
	slice1 := []int{1, 2, 3, 4, 5}
	slice2 := []int{4, 5, 6, 7, 8}
	intersecao, err := IntersecaoSlices(slice1, slice2)
	if err != nil {
		fmt.Println("Erro:", err)
	} else {
		fmt.Println("Interseção dos slices:", intersecao)
	}

	sliceVazio := []int{}
	resultadoVazio, errVazio := IntersecaoSlices(slice1, sliceVazio)
	if errVazio != nil {
		fmt.Println("Erro:", errVazio)
	} else {
		fmt.Println("Resultado:", resultadoVazio)
	}
}
