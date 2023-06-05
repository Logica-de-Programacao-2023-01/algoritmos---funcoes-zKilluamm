package main

import (
	"errors"
	"fmt"
	"math"
)

func VerificarPrimo(numero int) (bool, error) {
	if numero < 0 {
		return false, errors.New("Número negativo")
	}

	if numero < 2 {
		return false, nil
	}

	limite := int(math.Sqrt(float64(numero)))
	for i := 2; i <= limite; i++ {
		if numero%i == 0 {
			return false, nil
		}
	}

	return true, nil
}

func main() {
	numero := 13
	primo, err := VerificarPrimo(numero)
	if err != nil {
		fmt.Println("Erro:", err)
	} else {
		fmt.Printf("%d é primo: %t\n", numero, primo)
	}

	numeroNegativo := -7
	primoNegativo, errNegativo := VerificarPrimo(numeroNegativo)
	if errNegativo != nil {
		fmt.Println("Erro:", errNegativo)
	} else {
		fmt.Printf("%d é primo: %t\n", numeroNegativo, primoNegativo)
	}
}
