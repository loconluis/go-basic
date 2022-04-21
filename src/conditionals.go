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
}
