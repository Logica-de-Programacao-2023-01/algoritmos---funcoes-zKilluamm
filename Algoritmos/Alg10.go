package main

import (
	"errors"
	"fmt"
	"sort"
)

func OrdenarSlice(slice []int) ([]int, error) {
	if len(slice) == 0 {
		return nil, errors.New("Slice vazio")
	}

	sliceOrdenado := make([]int, len(slice))
	copy(sliceOrdenado, slice)
	sort.Ints(sliceOrdenado)

	return sliceOrdenado, nil
}

func main() {
	slice := []int{5, 2, 8, 1, 9, 3}
	sliceOrdenado, err := OrdenarSlice(slice)
	if err != nil {
		fmt.Println("Erro:", err)
	} else {
		fmt.Println("Slice ordenado:", sliceOrdenado)
	}

	sliceVazio := []int{}
	resultadoVazio, errVazio := OrdenarSlice(sliceVazio)
	if errVazio != nil {
		fmt.Println("Erro:", errVazio)
	} else {
		fmt.Println("Resultado:", resultadoVazio)
	}
}
