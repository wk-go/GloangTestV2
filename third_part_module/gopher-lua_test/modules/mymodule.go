package modules

import (
	"github.com/yuin/gopher-lua"
	"math/rand"
)

func Loader(L *lua.LState) int {
	// register functions to the table
	mod := L.SetFuncs(L.NewTable(), exports)
	// register other stuff
	L.SetField(mod, "name", lua.LString("mymodule"))

	// returns the module
	L.Push(mod)
	return 1
}

var exports = map[string]lua.LGFunction{
	"myfunc": myfunc,
}

func myfunc(L *lua.LState) int {
	v := rand.Intn(100)
	L.Push(lua.LNumber(v))
	L.Push(lua.LString("value2"))
	return 2 // 返回值数量
}
