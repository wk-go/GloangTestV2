package json

import (
	"encoding/json"
	"testing"
)

type Person struct {
	Name string
	Age  uint
}

// 测试从二维数组结构解析json数据
func TestUnmarshalFromTwoDimensionalArray(t *testing.T) {
	data := [][]Person{
		{{Name: "Sam", Age: 18}, {Name: "Jack", Age: 20}},
		{{Name: "LiLei", Age: 22}, {Name: "Joker", Age: 26}},
	}
	jsonData, err := json.Marshal(data)
	if err != nil {
		t.Error(err)
	}
	t.Logf("%s\n", jsonData)

	var data2 = make([][]Person, 0)
	err = json.Unmarshal(jsonData, &data2)
	if err != nil {
		t.Error(err)
	}
	t.Logf("%#v\n", data2)
}
