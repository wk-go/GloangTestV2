package main

import (
	"fmt"
	"github.com/yuin/gluamapper"
	"github.com/yuin/gopher-lua"
)

type Role struct {
	Name string
}

type Person struct {
	Name      string
	Age       int
	WorkPlace string
	Role      []*Role
}

func main() {

	L := lua.NewState()
	if err := L.DoString(`
person = {
  name = "Michel",
  age  = "31", -- weakly input
  work_place = "San Jose",
  role = {
    {
      name = "Administrator"
    },
    {
      name = "Operator"
    }
  }
}
`); err != nil {
		panic(err)
	}
	var person Person
	if err := gluamapper.Map(L.GetGlobal("person").(*lua.LTable), &person); err != nil {
		panic(err)
	}
	fmt.Printf("person info: %#v\n", person)
	for _, role := range person.Role {
		fmt.Printf("role of person: %#v\n", role)
	}
}
