package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

//Solicite ao usuário uma string e informe se ela contém pelo menos um número.

func main() {
	scaneador := bufio.NewScanner(os.Stdin)
	fmt.Println("informe sua string:")
	scaneador.Scan()
	nova_string := scaneador.Text()

	temNumero := false
	for _, r := range nova_string {
		if unicode.IsDigit(r) {
			temNumero = true
			break

		}

	}
	if temNumero == true {
		fmt.Println("a sua string contem um numero ou mais.")

	} else {
		fmt.Println("a sua string nao contem nenhum numero")
	}

}
