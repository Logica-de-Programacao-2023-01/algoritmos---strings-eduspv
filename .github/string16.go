package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

//Solicite ao usuário duas strings e informe se a segunda é uma substring da primeira.

func main() {
	s1 := bufio.NewScanner(os.Stdin)
	s2 := bufio.NewScanner(os.Stdin)
	fmt.Println("digite sua primeira string:")
	s1.Scan()
	fmt.Println("digite a sua segunda string:")
	s2.Scan()
	if strings.Contains(s1.Text(), s2.Text()) {
		fmt.Println("A sua segunda string é uma substring da primeira.")

	} else {

		fmt.Println("A sua segunda string não é uma substring da primeira.")
	}
}
