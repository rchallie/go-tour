package main

import "fmt"

const Pi = 3.14

func main() {
	const test = "world !"
	fmt.Println("Hello", test);
	fmt.Println("Happy", Pi, "Day !")

	const Truth = true
	fmt.Println("Go rules ?", Truth)
}
