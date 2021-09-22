package main

import (
	"fmt"
	"reflect"
)

var i, j int = 1, 2

func main() {
	var c, python, java = true, false, "no!"
	fmt.Println(i, j, c, python, java)
	fmt.Println(reflect.TypeOf(i), reflect.TypeOf(j), reflect.TypeOf(c), reflect.TypeOf(python), reflect.TypeOf(java))
}
