package main

import (
	"encoding/json"
	"fmt"
)

type Template struct {
	Code int64       `json:"code,omitempty"`
	Msg  string      `json:"msg,omitempty"`
	Data interface{} `json:"data"`
}

type Data1 struct {
	A int `json:"a"`
}

type Data2 struct {
	B string `json:"b"`
}

func main() {
	value1 := []byte(`{"code":123,"msg":"445","data":{"a":123}}`)
	value2 := []byte(`{"code":123,"msg":"445","data":{"b":"123"}}`)

	var d1 Data1
	var d2 Data2
	t1 := Template{Data: d1}
	t2 := Template{Data: d2}
	json.Unmarshal(value1, &t1)
	json.Unmarshal(value2, &t2)
	fmt.Printf("template\n%+v\n%+v\n", t1, t2)
	fmt.Printf("value\n%+v\n%+v\n", d1, d2)
	fmt.Println("---------")
	t1 = Template{Data: &d1}
	t2 = Template{Data: &d2}
	json.Unmarshal(value1, &t1)
	json.Unmarshal(value2, &t2)
	fmt.Printf("template\n%+v\n%+v\n", t1, t2)
	fmt.Printf("value\n%+v\n%+v\n", d1, d2)
}

/* output:

template
{Code:123 Msg:445 Data:map[a:123]}
{Code:123 Msg:445 Data:map[b:123]}
value
{A:0}
{B:}
---------
template
{Code:123 Msg:445 Data:0xc04206c0a0}
{Code:123 Msg:445 Data:0xc042056230}
value
{A:123}

*/
