package mypack

//go:generate go run gen/gen.go arg1 arg2
func PackFunc() string { return "mypack.PackFunc" }
