package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "string"
	s = strings.Replace(s, "s", "x", 1)
	fmt.Printf(s)
}
