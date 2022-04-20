package main

import "fmt"

func main() {
	// Area de un cuadrado
	const baseCuadrado = 10
	areaCuadrado := baseCuadrado * baseCuadrado
	fmt.Println("Area cuadrado: ", areaCuadrado)

	x := 10
	y := 50
	// Suma
	result := x + y
	fmt.Println("Suma:", result)

	// Resta
	result = y - x
	fmt.Println("Resta:", result)

	// Multiplicacion
	result = x * y
	fmt.Println("Multiplicacion:", result)

	// Division
	result = y / x
	fmt.Println("Division:", result)

	// Modulo / Residuo
	result = y % x
	fmt.Println("Modulo:", result)

	// Incremento
	x++
	fmt.Println("Incremental:", x)

	// Decremento
	x--
	fmt.Println("Incremental:", x)
	fmt.Println("****************** EJERCICIO ******************")
	// Area de rectangulo
	baseRectangulo := 15
	alturaRectangulo := 10

	areaRectangulo := baseRectangulo * alturaRectangulo
	fmt.Println("Area Rectangulo:", areaRectangulo)

	// Area de rectangulo
	baseTrapecio1 := 15
	baseTrapecio2 := 5
	alturaTrapecio := 10

	areaTrapecio := ((baseTrapecio1 + baseTrapecio2) / 2) * alturaTrapecio
	fmt.Println("Area Trapecio:", areaTrapecio)

	// Area de Circulo
	pi := 3.14159265358979323846
	const radio float64 = 15

	areaCirculo := pi * (radio * radio)
	fmt.Println("Area Circulo:", areaCirculo)
}
