package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func main() {
		v := Vertex{1, 2}
		fmt.Println(v)

		v.X = 4
		fmt.Println(v.X)

		p := &v

		p.X = 1e9
		// same as (*p).X

		fmt.Println(v, p.X, (*p).X)
}
