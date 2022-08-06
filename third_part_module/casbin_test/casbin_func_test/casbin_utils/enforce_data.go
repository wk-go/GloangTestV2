package casbin_utils

import (
	"github.com/casbin/casbin/v2"
	"gorm.io/gorm"
	"log"
)

var DB *gorm.DB

var Enforcer *casbin.Enforcer

func init() {
	e, err := casbin.NewEnforcer("./conf/model.conf", "./conf/policy.csv")
	if err != nil {
		log.Fatal(err)
	}
	Enforcer = e
}

type EnforceData struct {
	User    User
	Params  map[string]interface{}
	Data    []interface{} // 权限判断数据，顺序与Enforce函数一致
	Query   interface{}   // 前置查询条件，可以根据权限预设查询条件
	Subject string
	Object  string
	Action  string
	Veto    bool // 一票否决
}

func NewEnforceData(user User, sub, obj, act string, params map[string]interface{}) *EnforceData {
	return &EnforceData{
		User:    user,
		Subject: sub,
		Object:  obj,
		Action:  act,
		Params:  params,
		Veto:    true,
	}
}

func (d *EnforceData) GetData() []interface{} {
	d.Data = []interface{}{d.Subject, d.Object, d.Action, d}
	return d.Data
}

// SetQuery 这里实现了一个根据权限前置查询的操作
func (d *EnforceData) SetQuery(query interface{}, args ...interface{}) {
	if query == nil {
		return
	}
	if d.Query == nil {
		d.Query = DB.Where(query, args)
		return
	}
	if _q, ok := d.Query.(*gorm.DB); ok {
		_q.Where(query, args)
		return
	}
	d.Query = query
}
