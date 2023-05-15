package main

import (
	"fmt"
	"strings"
)

//Escreva um programa que receba uma string e converta todas as letras minúsculas para maiúsculas.
//Informe ao usuário o resultado.

func main() {
	string := "VASCO DA GAMA MEU BEM CAMPEÃO DO CEU E DO MAR"
	fmt.Println(string)
	string = strings.ToLower(string)
	fmt.Print(string)

}
