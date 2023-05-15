package main

import (
	"bufio"
	"fmt"
	"os"
)

//Escreva um programa que receba uma string e verifique se ela é um palíndromo.
//Informe ao usuário se é ou não.

func main() {
	DeTras := ""
	DeFrente := bufio.NewScanner(os.Stdin)
	fmt.Println("digite a sua string:")
	DeFrente.Scan()
	for i := len(DeFrente.Text()) - 1; i >= 0; i-- {
		DeTras += string(DeFrente.Text()[i])
	}
	if DeTras == DeFrente.Text() {
		fmt.Println("A sua string é um palíndromo!")

	} else {
		fmt.Println("a sua strimg não é um palíndromo.")

	}

}
