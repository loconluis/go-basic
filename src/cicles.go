package main

import "fmt"

func main() {
	// For conditional
	for i := 0; i < 10; i++ {
		fmt.Println("For:", i)
	}

	fmt.Println("************************************")

	// For while
	counter := 0
	for counter < 10 {
		fmt.Println("For While:", counter)
		counter++
	}

	fmt.Println("************************************")

	// For forever
	counterForever := 0
	for {
		fmt.Println("For Forever:", counterForever)
		counterForever++
		if counterForever == 10 {
			break
		}
	}
}
