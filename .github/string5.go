package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// Escreva um programa que receba uma string
// e verifique se ela é um número válido em ponto flutuante.
// Informe ao usuário se é ou não.
func main() {
	texto := bufio.NewScanner(os.Stdin)
	fmt.Println("Digite a sua string:")
	texto.Scan()
	if _, err := strconv.ParseFloat(texto.Text(), 64); err == nil {

		fmt.Println("Sua string é float")
	} else {
		fmt.Println("a sua string não é float")
	}
}
