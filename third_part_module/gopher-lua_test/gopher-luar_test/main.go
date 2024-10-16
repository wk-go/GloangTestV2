package main

import (
	"fmt"
	"github.com/yuin/gopher-lua"
	"layeh.com/gopher-luar"
)

func main() {
	Example_basic()
}

type User struct {
	Name  string
	token string
}

func (u *User) SetToken(t string) {
	u.token = t
}

func (u *User) Token() string {
	return u.token
}

const script = `
print("Hello from Lua, " .. u.Name .. "!" .. ", token is " .. u:Token())
u:SetToken("lua12345")
`

func Example_basic() {
	L := lua.NewState()
	defer L.Close()

	u := &User{
		Name:  "Tim",
		token: "go654321",
	}
	L.SetGlobal("u", luar.New(L, u))
	if err := L.DoString(script); err != nil {
		panic(err)
	}

	fmt.Println("Lua set your token to:", u.Token())
	// Output:
	// Hello from Lua, Tim!, token is go654321
	// Lua set your token to: 12345
}
