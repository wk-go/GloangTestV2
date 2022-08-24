package main

import "fmt"

// Struct Person has methods on both value and pointer receivers. Such usage is not recommended by the Go Documentation.

type Person struct {
	age int
}

func (p Person) howOld() int {
	return p.age
}

func (p *Person) growUp() {
	p.age++
}

func main() {
	//qcrao 是值类型
	qcrao := Person{age: 18}

	fmt.Println(qcrao.howOld())
	qcrao.growUp()
	fmt.Println(qcrao.howOld())

	//stefno 是指针类型
	stefno := &Person{age: 100}
	fmt.Println(stefno.howOld())
	stefno.growUp()
	fmt.Println(stefno.howOld())
}
