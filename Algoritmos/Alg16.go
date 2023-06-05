package main

import (
	"errors"
	"fmt"
	"strings"
)

func SubstituirVogaisPorAsterisco(str string) (string, error) {
	if len(str) == 0 {
		return "", errors.New("String vazia")
	}

	vogais := []string{"a", "e", "i", "o", "u", "A", "E", "I", "O", "U"}

	for _, vogal := range vogais {
		str = strings.ReplaceAll(str, vogal, "*")
	}

	return str, nil
}

func main() {
	str := "Hello World"
	novaStr, err := SubstituirVogaisPorAsterisco(str)
	if err != nil {
		fmt.Println("Erro:", err)
	} else {
		fmt.Println("Nova string:", novaStr)
	}

	strVazia := ""
	resultadoVazio, errVazio := SubstituirVogaisPorAsterisco(strVazia)
	if errVazio != nil {
		fmt.Println("Erro:", errVazio)
	} else {
		fmt.Println("Resultado:", resultadoVazio)
	}
}
