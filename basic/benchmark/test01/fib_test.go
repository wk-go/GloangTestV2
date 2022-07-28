package test01

import "testing"

var fibMap = map[int]int{
	1:  1,
	2:  1,
	3:  2,
	4:  3,
	5:  5,
	6:  8,
	7:  13,
	8:  21,
	9:  34,
	10: 55,
	11: 89,
}

func BenchmarkFib(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Fib(20) // 运行 Fib 函数 N 次
	}
}

func BenchmarkFib2(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Fib2(20) // 运行 Fib 函数 N 次
	}
}

func TestFib2(t *testing.T) {
	for k, v := range fibMap {
		_v := Fib2(k)
		if v != _v {
			t.Errorf("Number:%d; Expected:%d; Got:%d\n", k, v, _v)
		}
	}
}

func TestFib(t *testing.T) {
	for k, v := range fibMap {
		_v := Fib(k)
		if v != _v {
			t.Errorf("Number:%d; Expected:%d; Got:%d\n", k, v, _v)
		}
	}
}
