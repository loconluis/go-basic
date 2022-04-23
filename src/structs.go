package main

import "fmt"

type car struct {
	brand string
	model int
}

func main() {
	myCar := car{brand: "Ford", model: 2018}
	fmt.Println(myCar)

	// Otra manera
	var myCar2 car
	myCar2.brand = "Ford"
	fmt.Println(myCar2)
}
