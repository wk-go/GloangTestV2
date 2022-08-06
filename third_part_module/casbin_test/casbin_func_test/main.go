package main

import (
	"casbin_func_test/casbin_utils"
	"log"
)

func main() {
	enforceData := casbin_utils.NewEnforceData(nil, "admin", "/admin", "get", nil)
	ok, err := casbin_utils.Enforcer.Enforce(enforceData.GetData()...)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(ok)

	enforceData2 := casbin_utils.NewEnforceData(nil, "user_manager", "/admin/user", "get", nil)
	ok, err = casbin_utils.Enforcer.Enforce(enforceData2.GetData()...)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(ok)
}
