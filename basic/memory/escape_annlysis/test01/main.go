package main

import "fmt"

func foo() *int {
	i := 3
	return &i
}

func main() {
	x := foo()
	fmt.Println(x)
}
