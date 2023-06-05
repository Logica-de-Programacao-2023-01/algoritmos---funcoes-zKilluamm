package main

import (
	"fmt"
	"math/rand"
	"time"
)

func jogarJogo() int {
	resposta := rand.Intn(100) + 1
	tentativas := 0

	for {
		var valor int
		fmt.Print("Digite um número entre 1 e 100: ")
		fmt.Scanf("%d", &valor)
		tentativas++

		if valor == resposta {
			fmt.Println("Parabéns, você acertou!")
			fmt.Println("Você utilizou", tentativas, "tentativas.")
			break
		} else if valor < resposta {
			fmt.Println("O número é maior que", valor, ".")
		} else {
			fmt.Println("O número é menor que", valor, ".")
		}
	}

	return tentativas
}

func jogarNovamente() bool {
	var resposta string
	fmt.Print("Você deseja jogar novamente? (s/n): ")
	fmt.Scanf("%s", &resposta)
	return resposta == "s" || resposta == "S"
}

func exibirTodasAsJogadas(jogadas []int) {
	for i, tentativas := range jogadas {
		fmt.Printf("Você utilizou %d tentativas na %dª jogada.\n", tentativas, i+1)
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())

	fmt.Println("Bem-vindo ao jogo da adivinhação!")

	jogadas := []int{}

	for {
		tentativas := jogarJogo()
		jogadas = append(jogadas, tentativas)

		if !jogarNovamente() {
			break
		}
	}

	fmt.Println()
	exibirTodasAsJogadas(jogadas)
}
