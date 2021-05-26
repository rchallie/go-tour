package main

import "fmt"

func main() {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println("Sum:", sum)

	til := 0
	for ; til < 50; {
		til++
	}

	fmt.Println("Til:", til)

	whl := 0

	for whl < 50 {
		whl = whl * 2 + 1
	}

	fmt.Println("Whl:", whl)

	// Infinite loop:
	/*
	for {
	}
	*/
}
