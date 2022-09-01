package main

// 反射结构体

import (
	"fmt"
	"reflect"
)

type Person struct {
	Name   string
	Gender string
	Age    int
}

type Employee struct {
	Person
	Department string
	Position   string
}

func main() {
	var employee = Employee{
		Person:     Person{Name: "Sam", Gender: "male", Age: 25},
		Department: "IT", Position: "Developer",
	}
	fmt.Println("&employee", reflect.ValueOf(&employee))
	fmt.Println("&employee.Kind()", reflect.ValueOf(&employee).Kind())
	// 访问指针指向的数据
	fmt.Println("&employee.Elem()", reflect.ValueOf(&employee).Elem())
	fmt.Println("&employee.Elem().Kind", reflect.ValueOf(&employee).Elem().Kind())
	// 可以直接访问嵌套的结构体
	fmt.Println("&employee.Elem().FieldByName(\"Person\")", reflect.ValueOf(&employee).Elem().FieldByName("Person"))
	fmt.Println("&employee.Elem().FieldByName(\"Person\").Kind()", reflect.ValueOf(&employee).Elem().FieldByName("Person").Kind())
	// 还可以直接访问嵌套结构体的属性
	fmt.Println("&employee.Elem().FieldByName(\"Name\")", reflect.ValueOf(&employee).Elem().FieldByName("Name"))
	fmt.Println("&employee.Elem().FieldByName(\"Name\").Kind()", reflect.ValueOf(&employee).Elem().FieldByName("Name").Kind())
	// 正常访问属性
	fmt.Println("&employee.Elem().FieldByName(\"Department\")", reflect.ValueOf(&employee).Elem().FieldByName("Department"))
	fmt.Println("&employee.Elem().FieldByName(\"Department\").Kind()", reflect.ValueOf(&employee).Elem().FieldByName("Department").Kind())

	// 这样访问会报错: panic: reflect: call of reflect.Value.FieldByName on ptr Value
	fmt.Println("&employee", reflect.ValueOf(&employee).FieldByName("Department"))

}
