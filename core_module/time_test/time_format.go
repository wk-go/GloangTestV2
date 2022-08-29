package main

import (
	"fmt"
	"time"
)

func main() {
	s := "2021-6-1 14:20:25"
	x, err := time.Parse("2006-1-2 15:04:05", s)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(x)

	s = ""
	x, err = time.Parse("2006-1-2 15:04:05", s)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(x)

	b := []byte("\"\"")
	x, err = time.Parse("2006-1-2", string(b[1:len(b)-1]))
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(x)

}
