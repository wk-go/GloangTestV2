package main

import (
	"fmt"
	"os"
)

// param demo: $v1+$v2+$v3*$v4-($v5-$v6)/$v7+$v8 92 5 5 27 92 12 4 26
func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		args = []string{"$v1+$v2+$v3*$v4-($v5-$v6)/$v7+$v8", "92", "5", "5", "27", "92", "12", "4", "26"}
	}
	fmt.Println("args:", args)
	tidyStr := ExpressionTidy(args[0], args[1:]...)
	fmt.Println("tidyStr:", tidyStr)
	expression := args[0]
	numbers := make([]any, len(args[1:]))
	for i, v := range args[1:] {
		numbers[i] = v
	}
	value, err := Calculate(expression, numbers...)
	if err != nil {
		fmt.Println("err:", err)
		return
	}
	fmt.Println("value:", value)
}
