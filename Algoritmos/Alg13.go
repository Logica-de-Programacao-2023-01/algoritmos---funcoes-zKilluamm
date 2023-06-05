package main

import (
	"errors"
	"fmt"
)

func SomarDigitos(numero int) (int, error) {
	if numero < 0 {
		return 0, errors.New("Número negativo")
	}

	soma := 0
	for numero != 0 {
		digito := numero % 10
		soma += digito
		numero /= 10
	}

	return soma, nil
}

func main() {
	numero := 12345
	soma, err := SomarDigitos(numero)
	if err != nil {
		fmt.Println("Erro:", err)
	} else {
		fmt.Printf("A soma dos dígitos de %d é: %d\n", numero, soma)
	}

	numeroNegativo := -9876
	somaNegativa, errNegativa := SomarDigitos(numeroNegativo)
	if errNegativa != nil {
		fmt.Println("Erro:", errNegativa)
	} else {
		fmt.Printf("A soma dos dígitos de %d é: %d\n", numeroNegativo, somaNegativa)
	}
}
