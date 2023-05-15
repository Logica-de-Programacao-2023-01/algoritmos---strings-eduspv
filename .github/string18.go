package main

import (
	"bufio"
	"fmt"
	"os"
)

//Solicite ao usuário uma string e informe se ela contém somente números (0 a 9).

func main() {
	s1 := bufio.NewScanner(os.Stdin)
	fmt.Println("digite uma string:")
	s1.Scan()
	verify := s1.Text()
	for _, char := range verify {
		if char < '0' || char > '9' {
			fmt.Println("a sua string não contém somente numeros")
			return
		}
	}

	fmt.Println("A sua string contém somente numeros:")
}
