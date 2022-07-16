package json

import (
	"encoding/json"
	"fmt"
)

type Monster struct {
	//对于结构体的序列化，如果希望序列化后 key 的名字为自定义的名称，那么可以给 struct 字段指定一个 tag 标签
	//使用反引号，即波浪号下面的符号重命名。
	//并且由于序列化是使用的外部包（json.Marshal(...)），所以结构体的字段名首字母就必须是大写的，否则无权限。
	Name  string `json:"monster_name"` //反射机制
	Age   int    `json:"monster_age"`
	Skill string
}

func structJson() {
	monster := Monster{
		Name:  "牛魔王",
		Age:   500,
		Skill: "牛魔拳",
	}

	//将 monster 序列化
	data, err := json.Marshal(&monster)
	if err != nil {
		fmt.Printf("序列号错误 err=%v\n", err)
	}
	fmt.Printf("monster序列化后=%v\n", string(data))
}

func mapJson() {
	//定义一个 map，key 为 string 类型，value 为任一数据类型
	var a map[string]interface{}
	a = make(map[string]interface{})
	a["name"] = "红孩儿"
	a["age"] = 30
	a["address"] = "洪崖洞"

	data, err := json.Marshal(a)
	if err != nil {
		fmt.Printf("序列化错误 err=%v\n", err)
	}
	fmt.Printf("a map 序列化后=%v\n", string(data))
}

func sliceJson() {
	var slice []map[string]interface{}
	var m1 map[string]interface{}
	m1 = make(map[string]interface{})

	m1["name"] = "jack"
	m1["age"] = "7"
	m1["address"] = "北京"
	slice = append(slice, m1)

	var m2 map[string]interface{}
	m2 = make(map[string]interface{})
	m2["name"] = "tom"
	m2["age"] = "20"
	m2["address"] = [2]string{"墨西哥", "夏威夷"}
	slice = append(slice, m2)

	//将切片进行序列化操作
	data, err := json.Marshal(slice)
	if err != nil {
		fmt.Printf("序列化错误 err=%v\n", err)
	}
	fmt.Printf("slice 序列化后=%v\n", string(data))
}

func testFloat64() {
	var num1 float64 = 2345.67

	//对基本数据类型序列化，但对基本数据类型进行序列化意义不大，因为输出的还是基本数据类型值
	data, err := json.Marshal(num1)
	if err != nil {
		fmt.Printf("序列化错误 err=%v\n", err)
	}
	fmt.Printf("num1 序列化后=%v\n", string(data))
}
