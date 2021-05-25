package main

import "fmt"

func add_and_less(a , b int) (int , int) {
	return a + b, a - b
}

func main() {
	a, b := add_and_less(42, 24)
	fmt.Println(a, b)
}
