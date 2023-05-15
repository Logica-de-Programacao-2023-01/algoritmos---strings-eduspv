package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

//Solicite ao usuário uma string e substitua todas as ocorrências
//de uma letra por outra informada pelo usuário.

func main() {
	texto := bufio.NewScanner(os.Stdin)
	sair := bufio.NewScanner(os.Stdin)
	entrar := bufio.NewScanner(os.Stdin)
	fmt.Println("digite uma string:")
	texto.Scan()
	fmt.Println("digite a letra que voce quer que seja substituida:")
	sair.Scan()
	fmt.Println("digite a letra que voce vai substituir:")
	entrar.Scan()
	NovaString := strings.ReplaceAll(texto.Text(), sair.Text(), entrar.Text())
	fmt.Println(NovaString)
}
