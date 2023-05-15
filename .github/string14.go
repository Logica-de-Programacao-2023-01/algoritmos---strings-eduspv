package main

import (
	"bufio"
	"fmt"
	"os"
)

//Solicite ao usuário uma string e informe se ela é uma sequência numérica decrescente (exemplo: "987" ou "54321").

func main() {
	scanneador := bufio.NewScanner(os.Stdin)
	fmt.Print("digite uma string de numeros:")
	scanneador.Scan()
	texto := scanneador.Text()
	if len(texto) <= 1 {
		fmt.Print("A sequência não é decrescente.")
		return
	}

	for i := len(texto) - 2; i >= 0; i-- {
		if texto[i] < texto[i+1]-1 {
			fmt.Println("A sequência não é decrescente")
			return
		}
	}

	fmt.Println("A sequência é decrescente")
}
