package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// Leitura da string informada pelo usuário
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Informe uma string: ")
	input, _ := reader.ReadString('\n')

	// Remover espaços em branco ao redor da string
	input = strings.TrimSpace(input)

	// Inicializa o contador
	count := 0

	// Percorre a string e conta quantas vezes a letra 'a' ou 'A' aparece
	for _, char := range input {
		if char == 'a' || char == 'A' {
			count++
		}
	}

	// Verifica se há ocorrência da letra 'a' ou 'A'
	if count > 0 {
		fmt.Printf("A letra 'a' ou 'A' aparece %d vez(es) na string.\n", count)
	} else {
		fmt.Println("A letra 'a' ou 'A' não foi encontrada na string.")
	}
}
