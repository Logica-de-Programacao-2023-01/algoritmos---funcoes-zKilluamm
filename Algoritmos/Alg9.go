package main

import (
	"errors"
	"fmt"
	"strings"
)

func ExtrairPalavras(str string) ([]string, error) {
	if str == "" {
		return nil, errors.New("String vazia")
	}

	palavras := strings.Fields(str)
	return palavras, nil
}

func main() {
	str := "Olá, mundo! Como vai você?"
	palavras, err := ExtrairPalavras(str)
	if err != nil {
		fmt.Println("Erro:", err)
	} else {
		fmt.Println("Palavras:", palavras)
	}

	strVazia := ""
	resultadoVazio, errVazio := ExtrairPalavras(strVazia)
	if errVazio != nil {
		fmt.Println("Erro:", errVazio)
	} else {
		fmt.Println("Resultado:", resultadoVazio)
	}
}
