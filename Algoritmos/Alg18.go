package main

import (
	"errors"
	"fmt"
)

func AplicarFuncaoEmSlice(slice []int, funcao func(int) int) (int, error) {
	if len(slice) == 0 {
		return 0, errors.New("Slice vazio")
	}

	soma := 0
	for _, num := range slice {
		resultado := funcao(num)
		soma += resultado
	}

	return soma, nil
}

func dobrarNumero(numero int) int {
	return numero * 2
}

func main() {
	slice := []int{1, 2, 3, 4, 5}
	resultado, err := AplicarFuncaoEmSlice(slice, dobrarNumero)
	if err != nil {
		fmt.Println("Erro:", err)
	} else {
		fmt.Println("Resultado:", resultado)
	}

	sliceVazio := []int{}
	resultadoVazio, errVazio := AplicarFuncaoEmSlice(sliceVazio, dobrarNumero)
	if errVazio != nil {
		fmt.Println("Erro:", errVazio)
	} else {
		fmt.Println("Resultado:", resultadoVazio)
	}
}
