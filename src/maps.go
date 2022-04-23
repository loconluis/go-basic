package main

import "fmt"

func main() {
	m := make(map[string]int)
	m["jose"] = 14
	m["Pepito"] = 15

	fmt.Println(m)

	// Recorrer un map
	for key, value := range m {
		fmt.Println(key, value)
	}

	// Encontrar un valor
	value, ok := m["jose"]
	fmt.Println("ok", ok)
	fmt.Println(value)
}
