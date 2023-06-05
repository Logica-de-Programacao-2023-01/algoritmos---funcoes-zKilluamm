package main

import (
	"errors"
	"fmt"
)

type FuncaoInt func(int) int

func AplicarFuncao(slice []int, funcao FuncaoInt) ([]int, error) {
	if len(slice) == 0 {
		return nil, errors.New("Slice vazio")
	}

	resultado := make([]int, len(slice))
	for i, num := range slice {
		resultado[i] = funcao(num)
	}

	return resultado, nil
}

func main() {
	slice := []int{1, 2, 3, 4, 5}
	dobroSlice, err := AplicarFuncao(slice, func(x int) int {
		return x * 2
	})
	if err != nil {
		fmt.Println("Erro:", err)
	} else {
		fmt.Println("Dobro do slice:", dobroSlice)
	}

	sliceVazio := []int{}
	resultadoVazio, errVazio := AplicarFuncao(sliceVazio, func(x int) int {
		return x * 2
	})
	if errVazio != nil {
		fmt.Println("Erro:", errVazio)
	} else {
		fmt.Println("Resultado:", resultadoVazio)
	}
}
