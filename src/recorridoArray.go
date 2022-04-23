package main

import (
	"fmt"
	"strings"
)

func isPalindromo(text string) {
	parseText := strings.ToLower(text)
	var textReverse string

	for i := len(parseText) - 1; i >= 0; i-- {
		textReverse += string(parseText[i])
	}

	if parseText == textReverse {
		fmt.Println("Es palindromo")
	} else {
		fmt.Println("No es palindromo")
	}
}

func main() {
	// slice := []string{"Hola", "como", "estas"}
	// for i, value := range slice {
	// 	fmt.Println("Index:", i, "Value:", value)
	// }

	text := "Ama"
	isPalindromo(text)
}
