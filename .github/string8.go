package main

import (
	"bufio"
	"fmt"
	"os"
)

//Escreva um programa que receba uma string e inverta a ordem dos caracteres. Informe ao usuário o resultado.

func main() {
	s := bufio.NewScanner(os.Stdin)
	AoContrario := ""
	fmt.Print("Digite a sua string:")
	s.Scan()
	for i := len(s.Text()) - 1; i >= 0; i-- {
		AoContrario += string(s.Text()[i])
	}
	fmt.Println(AoContrario)
}
