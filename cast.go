package main

import (
	"fmt"
	"strconv"
)

func main() {
	var x int = 1
	xf := float64(x)

	var f float64 = 1.1
	fx := int(f)

	var s string = "11"
	sx, _ := strconv.Atoi(s)
	//si, err := strconv.Atoi(s)

	fmt.Printf("%T %v %f\n", xf, xf, xf)
	fmt.Printf("%T %v %d\n", fx, fx, fx)
	fmt.Printf("%T %v %d\n", sx, sx, sx)
}
