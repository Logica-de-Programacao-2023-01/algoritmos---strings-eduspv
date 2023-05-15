package main

import (
	"bufio"
	"fmt"
	"os"
)

//Solicite ao usuário uma string e imprima a mesma string invertida.

func main() {
	s1 := bufio.NewScanner(os.Stdin)
	fmt.Println("digite a sua string:")
	s1.Scan()

}
