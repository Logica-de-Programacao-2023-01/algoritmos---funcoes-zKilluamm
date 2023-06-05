package main

import (
	"errors"
	"fmt"
)

func isPrimo(numero int) bool {
	if numero < 2 {
		return false
	}

	for i := 2; i*i <= numero; i++ {
		if numero%i == 0 {
			return false
		}
	}

	return true
}

func NumerosPrimos(numero int) ([]int, error) {
	if numero < 0 {
		return nil, errors.New("Número negativo")
	}

	primos := []int{}
	for i := 2; i <= numero; i++ {
		if isPrimo(i) {
			primos = append(primos, i)
		}
	}

	return primos, nil
}

func main() {
	numero := 20
	primos, err := NumerosPrimos(numero)
	if err != nil {
		fmt.Println("Erro:", err)
	} else {
		fmt.Println("Números primos até", numero, ":", primos)
	}

	numeroNegativo := -5
	primosNegativos, errNegativos := NumerosPrimos(numeroNegativo)
	if errNegativos != nil {
		fmt.Println("Erro:", errNegativos)
	} else {
		fmt.Println("Números primos até", numeroNegativo, ":", primosNegativos)
	}
}
