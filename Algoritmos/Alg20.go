package main

import (
	"errors"
	"fmt"
)

func FiltrarStrings(slice []string) ([]string, error) {
	if len(slice) == 0 {
		return nil, errors.New("Slice vazio")
	}

	filtradas := []string{}
	for _, str := range slice {
		if len(str) > 5 {
			filtradas = append(filtradas, str)
		}
	}

	return filtradas, nil
}

func main() {
	slice := []string{"apple", "banana", "orange", "kiwi", "mango"}
	filtradas, err := FiltrarStrings(slice)
	if err != nil {
		fmt.Println("Erro:", err)
	} else {
		fmt.Println("Strings filtradas:", filtradas)
	}

	sliceVazio := []string{}
	filtradasVazias, errVazio := FiltrarStrings(sliceVazio)
	if errVazio != nil {
		fmt.Println("Erro:", errVazio)
	} else {
		fmt.Println("Strings filtradas:", filtradasVazias)
	}
}
