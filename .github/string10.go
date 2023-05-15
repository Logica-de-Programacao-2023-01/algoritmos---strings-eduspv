package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	var s1, s2 string
	fmt.Println("Digite a primeira string:")
	fmt.Scanln(&s1)
	fmt.Println("Digite a segunda string:")
	fmt.Scanln(&s2)

	s1 = strings.ToLower(strings.ReplaceAll(s1, " ", ""))
	s2 = strings.ToLower(strings.ReplaceAll(s2, " ", ""))

	if len(s1) != len(s2) {
		fmt.Println("As strings não são anagramas.")
		return
	}

	slice1 := []byte(s1)
	slice2 := []byte(s2)
	sort.Slice(slice1, func(i, j int) bool { return slice1[i] < slice1[j] })
	sort.Slice(slice2, func(i, j int) bool { return slice2[i] < slice2[j] })

	if string(slice1) == string(slice2) {
		fmt.Println("As strings são anagramas!")
	} else {
		fmt.Println("As strings não são anagramas.")
	}
}
