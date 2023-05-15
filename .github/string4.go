package main

import (
	"bufio"
	"fmt"
	"os"
)

//Solicite ao usuário duas strings e informe se elas são iguais ou diferentes.

func main() {
	primeira_string := bufio.NewScanner(os.Stdin)
	segunda_string := bufio.NewScanner(os.Stdin)
	fmt.Println("Digite o primeiro texto:")
	primeira_string.Scan()
	fmt.Println("digite o segundo texto:")
	segunda_string.Scan()
	if primeira_string.Text() == segunda_string.Text() {
		fmt.Println("as duas strings são iguais.")

	} else {
		fmt.Print("as duas strings são diferentes.")
	}

}
