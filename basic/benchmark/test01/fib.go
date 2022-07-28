package test01

func Fib(n int) int {
	switch n {
	case 0:
		return 0
	case 1:
		return 1
	default:
		return Fib(n-1) + Fib(n-2)
	}
}

func Fib2(n int) int {
	i, j := 1, 0
	for m := 0; m < n; m++ {
		i, j = j, i+j
	}
	return j
}
