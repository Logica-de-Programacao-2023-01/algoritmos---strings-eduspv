package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

//Escreva um programa que receba uma string
//e conte quantas palavras ela contém.
//Informe ao usuário o resultado.

func main() {
	fmt.Println("digite sua string:")
	scaneador := bufio.NewScanner(os.Stdin)
	scaneador.Scan()
	texto := scaneador.Text()
	palavras := strings.Fields(texto)
	num_palavras := len(palavras)

	fmt.Printf("A string contém %d palavras.\n", num_palavras)
}
