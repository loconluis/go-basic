package main

import "fmt"

func normalFunction(message string) {
	fmt.Println(message)
}

func tripleArgument(a, b int, c string) {
	fmt.Println(a, b, c)
}

func returnOneValue(a int) int {
	return a * 2
}

func returnMultipleValues(a int) (int, int) {
	return a * 2, a * 3
}

func areaCuadrado(base int) int {
	return base * base
}

func areaCirculo(radio float64) float64 {
	const pi float64 = 3.14159265358979323846
	return pi * (radio * radio)
}

func areaTrapecios(base1, base2, altura int) int {
	return ((base1 + base2) / 2) * altura
}

func areaRectangulo(base, altura int) int {
	return base * altura
}

func main() {
	normalFunction("Hello World 2")
	tripleArgument(1, 2, "Hola")
	value := returnOneValue(5)
	fmt.Println("Value: ", value)
	value1, value2 := returnMultipleValues(5)
	// Scape values use underscore
	// value1, _ := returnMultipleValues(5)
	fmt.Println("Value1: ", value1)
	fmt.Println("Value2: ", value2)
	fmt.Println("****************** EJERCICIO ******************")

	_areaCuadrado := areaCuadrado(10)
	fmt.Println("Area cuadrado: ", _areaCuadrado)
	_areaCirculo := areaCirculo(15)
	fmt.Println("Area circulo: ", _areaCirculo)
	_areaTrapecio := areaTrapecios(15, 5, 10)
	fmt.Println("Area trapecio: ", _areaTrapecio)
	_areaRectangulo := areaRectangulo(15, 10)
	fmt.Println("Area rectangulo: ", _areaRectangulo)
}
