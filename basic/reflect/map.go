package main

// 反射map类型数据
import (
	"fmt"
	"reflect"
)

func main() {
	m := map[string]any{
		"key1": "val1",
		"key2": 1,
		"key3": 3.01,
	}

	rv := reflect.ValueOf(m)
	for _, v := range rv.MapKeys() {
		value := rv.MapIndex(v)
		fmt.Printf("Key_type:%s, Key_value:%s, Value_type:%s, Value_value:%v\n", v.Kind(), v, value.Kind(), value)
	}
}
