package main

import (
	"errors"
	"fmt"
	"strings"
	"unicode"
)

func FiltrarStringsMaiusculas(slice []string) (string, error) {
	if len(slice) == 0 {
		return "", errors.New("Slice vazio")
	}

	var result strings.Builder
	for _, str := range slice {
		if len(str) > 0 && unicode.IsUpper(rune(str[0])) {
			result.WriteString(str)
			result.WriteString(" ")
		}
	}

	return strings.TrimSpace(result.String()), nil
}

func main() {
	slice := []string{"Hello", "world", "OpenAI", "GPT-3"}
	stringsMaiusculas, err := FiltrarStringsMaiusculas(slice)
	if err != nil {
		fmt.Println("Erro:", err)
	} else {
		fmt.Println("Strings maiúsculas:", stringsMaiusculas)
	}

	sliceVazio := []string{}
	resultadoVazio, errVazio := FiltrarStringsMaiusculas(sliceVazio)
	if errVazio != nil {
		fmt.Println("Erro:", errVazio)
	} else {
		fmt.Println("Resultado:", resultadoVazio)
	}
}
