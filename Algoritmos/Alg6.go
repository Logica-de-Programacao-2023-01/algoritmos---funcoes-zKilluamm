package main

import (
	"errors"
	"fmt"
	"strings"
)

func ConcatenarStrings(slice []string) (string, error) {
	if len(slice) == 0 {
		return "", errors.New("Slice vazio")
	}

	resultado := strings.Join(slice, ",")
	return resultado, nil
}

func main() {
	slice := []string{"Maçã", "Banana", "Laranja"}
	resultado, err := ConcatenarStrings(slice)
	if err != nil {
		fmt.Println("Erro:", err)
	} else {
		fmt.Println("Resultado:", resultado)
	}

	sliceVazio := []string{}
	resultadoVazio, errVazio := ConcatenarStrings(sliceVazio)
	if errVazio != nil {
		fmt.Println("Erro:", errVazio)
	} else {
		fmt.Println("Resultado:", resultadoVazio)
	}
}
