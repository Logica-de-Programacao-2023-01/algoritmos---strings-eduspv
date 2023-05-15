//Solicite ao usuário uma string e imprima somente as suas letras únicas (que aparecem apenas uma vez).

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Digite uma string: ")
	scanner.Scan()
	input := scanner.Text()

	counts := make(map[rune]int)
	for _, char := range input {
		counts[char]++
	}

	uniqueLetters := ""
	for _, char := range input {
		if counts[char] == 1 {
			uniqueLetters += string(char)
		}
	}

	fmt.Println("Letras únicas:", uniqueLetters)
}
