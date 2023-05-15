package main

import (
	"fmt"
	"strings"
)

//Escreva um programa que receba uma string e remova
//todos os espaços em branco. Informe ao usuário o resultado.

func main() {
	string := "Bagunçou a minha vida mudou todos os meus planos"
	fmt.Println("a sua antiga string é: ", string)
	string = strings.ReplaceAll(string, " ", "")
	fmt.Print("A sua nova string é:", string)

}
