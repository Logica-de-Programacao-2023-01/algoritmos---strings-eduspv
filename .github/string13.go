package main

import (
	"bufio"
	"fmt"
	"os"
)

//Solicite ao usuário uma string e informe se ela é uma sequência numérica crescente (exemplo: "123" ou "4567").

func main() {
	scanneador := bufio.NewScanner(os.Stdin)
	fmt.Print("digite uma string de numeros:")
	scanneador.Scan()
	texto := scanneador.Text()
	if len(texto) <= 1 {
		fmt.Print("A sequência não é crescente.")
		return
	}

	for i := 1; i < len(texto); i++ {
		if texto[i] != texto[i-1]+1 {
			fmt.Println("A sequência não é crescente")
			return
		}
	}

	fmt.Println("A sequência é crescente")
}
