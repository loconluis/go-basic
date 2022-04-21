package main

import "fmt"

func main() {
	value1 := 2
	value2 := 2

	if value1 < value2 {
		fmt.Println("Value1 is less than Value2")
	} else if value1 > value2 {
		fmt.Println("Value1 is greater than Value2")
	} else {
		fmt.Println("Value1 is equal to Value2")
	}

	// Switch
	modulo := 5 % 2
	switch modulo {
	case 0:
		fmt.Println("Modulo is 0 and is par")
	default:
		fmt.Printf("Modulo is %d and is impar\n", modulo)
	}
}
