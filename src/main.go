package main

import "fmt"

func main() {
	// Declaracion de constantes
	const pi float64 = 3.14159265358979323846
	const pi2 = 3.1416
	fmt.Println("pi: ", pi)
	fmt.Println("pi2: ", pi2)

	// Declaracion de variables enteras
	base := 12
	var altura int = 24
	var area int
	fmt.Println(base, altura, area)

	//Zero values
	//Go asigna vaalores a variables vacías
	var a int     // 0
	var b float64 // 0
	var c string  // " " blank space
	var d bool    // false
	fmt.Println(a, b, c, d)

	// Area de un cuadrado
	const baseCuadrado = 10
	areaCuadrado := baseCuadrado * baseCuadrado
	fmt.Println("Area cuadrado: ", areaCuadrado)
}
