package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

//Solicite ao usuário uma string e substitua todas as vogais por '*' (asterisco).

func main() {
	string := bufio.NewScanner(os.Stdin)
	fmt.Print("digite uma string:")
	string.Scan()
	Outrastring := string.Text()
	Outrastring = strings.ReplaceAll(Outrastring, "A", "*")
	Outrastring = strings.ReplaceAll(Outrastring, "E", "*")
	Outrastring = strings.ReplaceAll(Outrastring, "I", "*")
	Outrastring = strings.ReplaceAll(Outrastring, "O", "*")
	Outrastring = strings.ReplaceAll(Outrastring, "U", "*")
	Outrastring = strings.ReplaceAll(Outrastring, "a", "*")
	Outrastring = strings.ReplaceAll(Outrastring, "e", "*")
	Outrastring = strings.ReplaceAll(Outrastring, "i", "*")
	Outrastring = strings.ReplaceAll(Outrastring, "o", "*")
	Outrastring = strings.ReplaceAll(Outrastring, "u", "*")
	fmt.Print(Outrastring)

}
