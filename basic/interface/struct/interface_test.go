package _struct

import (
	"fmt"
	"testing"
)

// 接口函数的另外一种调用方法
type S struct {
}

func (*S) Say(message string) {
	fmt.Println("Hello,", message)
}

type I interface {
	Say(string)
}

func TestInterface(t *testing.T) {
	I.Say(new(S), "World")
}
