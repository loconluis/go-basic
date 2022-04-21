package main

import "fmt"

func main() {
	helloMessage := "Hello"
	worldMessage := "Hello"
	// Imprime y salta la linea
	fmt.Println(helloMessage, worldMessage)
	fmt.Println(helloMessage, worldMessage)

	// printF
	nombre := "Platzi"
	cursos := 500
	fmt.Printf("%s, tiene mas de %d cursos\n", nombre, cursos)
	fmt.Printf("%v, tiene mas de %v cursos\n", nombre, cursos)

	// Sprint F va a generar un string pero no lo imprime en consola
	message := fmt.Sprintf("%s, tiene mas de %d cursos\n", nombre, cursos)
	fmt.Println(message)

	// Tipo de dato
	fmt.Printf("helloMessage: %T\n", helloMessage)
	fmt.Printf("cursos: %T \n", cursos)
}
