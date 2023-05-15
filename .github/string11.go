package main

import (
	"fmt"
	"strings"
)

//Escreva um programa que receba uma string e remova todas as vogais. Informe ao usuário o resultado.

func main() {
	s := "O rato roeu a roupa do rei de roma"
	s = strings.ReplaceAll(s, "a", "")
	s = strings.ReplaceAll(s, "A", "")
	s = strings.ReplaceAll(s, "E", "")
	s = strings.ReplaceAll(s, "I", "")
	s = strings.ReplaceAll(s, "O", "")
	s = strings.ReplaceAll(s, "U", "")
	s = strings.ReplaceAll(s, "a", "")
	s = strings.ReplaceAll(s, "e", "")
	s = strings.ReplaceAll(s, "i", "")
	s = strings.ReplaceAll(s, "o", "")
	s = strings.ReplaceAll(s, "u", "")
	fmt.Println(s)

}
