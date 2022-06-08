package main

import (
	"fmt"

	"golang.org/x/example/stringutil"
)

func main() {
	in := "Hello, OTUS!"
	out := stringutil.Reverse(in)
	fmt.Println(out)
}
