package main

import (
	"errors"
	"fmt"
)

func VerificarPresenca(numero int, slice []int) (bool, error) {
	if len(slice) == 0 {
		return false, errors.New("Slice vazio")
	}

	for _, num := range slice {
		if num == numero {
			return true, nil
		}
	}

	return false, nil
}

func main() {
	numero := 5
	slice := []int{1, 2, 3, 4, 5}
	presente, err := VerificarPresenca(numero, slice)
	if err != nil {
		fmt.Println("Erro:", err)
	} else {
		fmt.Printf("O número %d está presente no slice: %t\n", numero, presente)
	}

	sliceVazio := []int{}
	resultadoVazio, errVazio := VerificarPresenca(numero, sliceVazio)
	if errVazio != nil {
		fmt.Println("Erro:", errVazio)
	} else {
		fmt.Println("Resultado:", resultadoVazio)
	}
}
