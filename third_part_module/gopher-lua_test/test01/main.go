package main

import (
	"fmt"
	lua "github.com/yuin/gopher-lua"
	"gopher_lua_test/modules"
)

func main() {

	// create a new state
	L := lua.NewState()
	defer L.Close()
	if err := L.DoString(`
print("hello")
x = 1 +2
print("x:", x)
y = "hello world"
print("y:", y)
return y
`); err != nil {
		panic(err)
	}
	// the value of x is still available
	if err := L.DoString(`print("x still existed:",x)`); err != nil {
		panic(err)
	}

	L02 := lua.NewState()
	defer L02.Close()
	if err := L02.DoFile("./script/hello.lua"); err != nil {
		panic(err)
	}

	// bind a function
	L03 := lua.NewState()
	defer L03.Close()
	L03.SetGlobal("double", L03.NewFunction(Double))
	if err := L03.DoString(`print("double(10):", double(10))`); err != nil {
		panic(err)
	}

	// You can test an object type in Go way(type assertion) or using a Type() value.
	L04 := lua.NewState()
	defer L04.Close()
	if err := L04.DoString(`
a = "hello world"
print("The type of variable a: ", type(a))
return a
`); err != nil {
		panic(err)
	}
	lv := L04.Get(-1) // get the value at the top of the stack
	if str, ok := lv.(lua.LString); ok {
		// lv is LString
		fmt.Println("lv value", string(str))
	}
	if lv.Type() != lua.LTString {
		fmt.Printf("not a string: %v\n", lv.Type())
	}

	L05 := lua.NewState()
	defer L05.Close()
	if err := L05.DoString(`
t = {"hello", "world"}
print("The type of variable t: ",type(t))
return t
`); err != nil {
		panic(err)
	}

	lv2 := L05.Get(-1) // get the value at the top of the stack
	if tbl, ok := lv2.(*lua.LTable); ok {
		// lv is LTable
		fmt.Println("The Length Of t:", L.ObjLen(tbl))
	}

	// call a function from lua
	fmt.Println("\n\n---- call function from lua: L06 ----")
	L06 := lua.NewState()
	defer L06.Close()
	if err := L06.DoFile("./script/function.lua"); err != nil {
		panic(err)
	}

	fmt.Println("---- call fibonacci function ----")
	if err := L06.CallByParam(lua.P{
		Fn:      L06.GetGlobal("fibonacci"),
		NRet:    1,
		Protect: true,
	}, lua.LNumber(6)); err != nil {
		fmt.Printf("call function error: %s\n", err)
	}
	fmt.Printf("Stack top: %d\n", L06.GetTop())
	ret := L06.Get(-1) // returned value
	L06.Pop(1)         // remove received value
	fmt.Println("fibonacci(6):", ret)
	fmt.Printf("Stack top: %d\n", L06.GetTop())

	fmt.Println("---- call say_hello function ----")
	fn := L06.GetGlobal("say_hello")
	if err := L06.CallByParam(lua.P{Fn: fn, NRet: 0, Protect: true}, lua.LString("everyone")); err != nil {
		fmt.Printf("call function error: %s\n", err)
	}

	// coroutine call test
	fmt.Println("\n\n---- coroutine call test ----")
	co03, _ := L06.NewThread()
	coroFunc := L06.GetGlobal("coro").(*lua.LFunction)
	for {
		st, err, values := L06.Resume(co03, coroFunc, lua.LNumber(20))
		if st == lua.ResumeError {
			fmt.Println("function call yield break(error):", err.Error())
			break
		}

		for i, lv := range values {
			fmt.Printf("function call result: %v : %v\n", i, lv)
		}

		if st == lua.ResumeOK {
			fmt.Println("function call yield break(ok)")
			break
		}
	}

	// Opening a subset of builtin modules
	L07 := lua.NewState(lua.Options{SkipOpenLibs: true})
	defer L07.Close()
	for _, pair := range []struct {
		n string
		f lua.LGFunction
	}{
		{lua.LoadLibName, lua.OpenPackage},
		{lua.BaseLibName, lua.OpenBase},
		{lua.TabLibName, lua.OpenTable},
		{lua.OsLibName, lua.OpenOs},
	} {
		if err := L07.CallByParam(lua.P{
			Fn:      L07.NewFunction(pair.f),
			NRet:    0,
			Protect: true,
		}, lua.LString(pair.n)); err != nil {
			fmt.Printf("L07 error: %s", err)
		}
	}
	if err := L07.DoFile("script/subset/main.lua"); err != nil {
		fmt.Printf("L07 DoFile error: %s", err)
	}

	// Creating a module by Go
	fmt.Println("\n\n---- Creating a module by Go: L08 ----")
	L08 := lua.NewState()
	defer L08.Close()
	L08.PreloadModule("mymodule", modules.Loader)
	if err := L08.DoFile("script/go_module_test/main.lua"); err != nil {
		panic(err)
	}
}

func Double(L *lua.LState) int {
	lv := L.ToInt(1)            /* get argument */
	L.Push(lua.LNumber(lv * 2)) /* push result */
	return 1                    /* number of results */
}
