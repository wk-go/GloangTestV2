// Lua脚本路径问题

package main

import (
	lua "github.com/yuin/gopher-lua"
	"os"
	"path"
	"path/filepath"
)

var RootPath string

func main() {
	var err error
	L := lua.NewState()
	defer L.Close()
	RootPath, err = filepath.Abs(path.Dir(os.Args[0]))
	if err != nil {
		panic(err)
	}
	// 给虚拟机添加GetLuaPath方法
	L.SetGlobal("GetLuaPath", L.NewFunction(GetLuaPath))
	if err := L.DoFile("scripts/main.lua"); err != nil {
		panic(err)
	}

}

// GetLuaPath 在lua脚本中设置加载路径
func GetLuaPath(L *lua.LState) int {
	// 绝对路径
	L.Push(lua.LString(RootPath + "/scripts/functions/?.lua"))
	return 1
}
