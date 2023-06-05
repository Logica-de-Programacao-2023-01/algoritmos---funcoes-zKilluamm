package main

import (
	"errors"
	"fmt"
)

func FiltrarNumerosPares(slice []int) ([]int, error) {
	if len(slice) == 0 {
		return nil, errors.New("Slice vazio")
	}

	numerosPares := make([]int, 0)
	for _, num := range slice {
		if num%2 == 0 {
			numerosPares = append(numerosPares, num)
		}
	}

	return numerosPares, nil
}

func main() {
	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	paresSlice, err := FiltrarNumerosPares(slice)
	if err != nil {
		fmt.Println("Erro:", err)
	} else {
		fmt.Println("Números pares:", paresSlice)
	}

	sliceVazio := []int{}
	resultadoVazio, errVazio := FiltrarNumerosPares(sliceVazio)
	if errVazio != nil {
		fmt.Println("Erro:", errVazio)
	} else {
		fmt.Println("Resultado:", resultadoVazio)
	}
}
