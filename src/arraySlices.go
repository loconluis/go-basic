package main

import "fmt"

func main() {
	// array
	var array [4]int
	array[0] = 1
	array[1] = 2
	fmt.Println("Array:", array)
	fmt.Println("Cuantos elementos hay en Array:", len(array))
	fmt.Println("Capacidad de elementos dentro de Array:", cap(array))

	// Slices
	slice := []int{1, 2, 3, 4, 5, 6, 7}
	fmt.Println("Slice:", slice)
	fmt.Println("Cuantos elementos hay en slice:", len(slice))
	fmt.Println("Capacidad de elementos dentro de slice:", cap(slice))

	// Metodos en slices
	fmt.Println(slice[0])
	fmt.Println(slice[:3])
	fmt.Println(slice[2:4])
	fmt.Println(slice[4:])

	// Append
	slice = append(slice, 8)
	fmt.Println("Slice:", slice)

	// Append nueva lista
	newSlice := []int{9, 10, 11, 12}
	slice = append(slice, newSlice...)
	fmt.Println("Slice:", slice)
}
