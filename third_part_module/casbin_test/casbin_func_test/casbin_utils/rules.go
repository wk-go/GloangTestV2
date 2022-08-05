package casbin_utils

import (
	"errors"
	"strconv"
	"strings"
)

var AccessRuleList map[string]AccessRule

func init() {
	AccessRuleList = map[string]AccessRule{
		"owner": AccessRuleFunc(func(data *EnforceData) (bool, error) { return true, nil }),
	}
	//注册规则函数
	Enforcer.AddFunction("ExecRules", ExecRules)
}

type AccessRule interface {
	Run(data *EnforceData) (bool, error)
}

type AccessRuleFunc func(data *EnforceData) (bool, error)

func (f AccessRuleFunc) Run(data *EnforceData) (bool, error) {
	return f(data)
}

func ExecRules(args ...interface{}) (interface{}, error) {
	if len(args) < 2 {
		return false, errors.New("parameters length expected 2 got " + strconv.Itoa(len(args)))
	}
	rulesString := args[1].(string)
	if rulesString == "" {
		return true, nil
	}
	enforceData := args[0].(*EnforceData)
	rulesSlice := strings.Split(rulesString, "|")
	for _, ruleName := range rulesSlice {
		if rule, ok := AccessRuleList[ruleName]; ok {
			if result, err := rule.Run(enforceData); err != nil || result == false {
				return result, err
			}
		}
	}
	return true, nil
}
