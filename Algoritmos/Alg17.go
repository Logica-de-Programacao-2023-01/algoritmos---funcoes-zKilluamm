package main

import (
	"errors"
	"fmt"
	"sort"
	"strings"
)

func OrdenarStrings(slice []string) (string, error) {
	if len(slice) == 0 {
		return "", errors.New("Slice vazio")
	}

	sort.Strings(slice)
	novaStr := strings.Join(slice, " ")

	return novaStr, nil
}

func main() {
	slice := []string{"banana", "abacaxi", "laranja", "maçã"}
	novaStr, err := OrdenarStrings(slice)
	if err != nil {
		fmt.Println("Erro:", err)
	} else {
		fmt.Println("Nova string:", novaStr)
	}

	sliceVazio := []string{}
	resultadoVazio, errVazio := OrdenarStrings(sliceVazio)
	if errVazio != nil {
		fmt.Println("Erro:", errVazio)
	} else {
		fmt.Println("Resultado:", resultadoVazio)
	}
}
