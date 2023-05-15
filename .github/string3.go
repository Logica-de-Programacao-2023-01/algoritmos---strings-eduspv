package main

import (
	"fmt"
	"strings"
)

//Escreva um programa que receba uma string
//e substitua todas as ocorrências desse caractere na string
//por outro caractere especificado pelo usuário.
//Informe ao usuário o resultado.

func main() {
	var antiga_letra, nova_letra string
	string := "A sua mãe vai me amar!!!"
	fmt.Println(string)
	fmt.Println("Qual letra voce quer que seja substituido?")
	fmt.Scan(&antiga_letra)
	fmt.Println("por qual voce quer que troque?")
	fmt.Scan(&nova_letra)
	string = strings.ReplaceAll(string, antiga_letra, nova_letra)
	fmt.Print(string)

}
