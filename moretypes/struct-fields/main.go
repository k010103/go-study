package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func main() {
	v := Vertex{}
	fmt.Println(v.X)
	v.X = 4
	fmt.Println(v.X)
}
