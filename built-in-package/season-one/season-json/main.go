package main

import (
	"encoding/json"
	"fmt"
	"reflect"
	"time"
)

/*
1. go 类型和 json 相互转换
2. 结构体和 json 相互转换
3. 自定义的 json 格式
*/
func main() {
	fmt.Println(string(boolToJson(1 == 1)))
	fmt.Println(jsonToBool([]byte(`true`)))

	fmt.Println(string(intToJson(12)))

	boolFalse, _ := json.Marshal(false)
	fmt.Println(string(boolFalse))

	IntNumber, _ := json.Marshal(123)
	fmt.Println(string(IntNumber))

	var info Info
	info = Info{
		Name:    "XieWei",
		Age:     100,
		City:    "shangHai",
		Company: "Only Me",
	}
	fmt.Println(string(info.MarshalOp()))

	var otherInfo Info
	otherInfo.Name = ""
	otherInfo.Age = 20
	otherInfo.City = "BeiJing"
	otherInfo.Company = "Only You"
	fmt.Println(string(otherInfo.MarshalOp()))

	fmt.Println(UnMarshalExample([]byte(`{"name":"xieWei", "age": "20", "city_shanghai": "GuangDong"}`)))

	var self = SelfMarshal{}
	self.Age = 20
	self.Name = "XieWei"
	self.City = "HangZhou"

	selfJsonMarshal, err := json.Marshal(self)
	fmt.Println(err, string(selfJsonMarshal))

	var t = Test{}
	t.Date = jsonTime(time.Now())
	t.Name = "Hello World"
	t.State = true
	body, _ := json.Marshal(t)
	fmt.Println(string(body))

	//var name = make(map[string]interface{})
	//name["key1"] = 123
	//name["key2"] = "string key"
	//name["key3"] = []string{"How", "Are", "You"}
	//name["key4"] = info
	var name = []byte(`{"key1":123,"key2":"string key", "key3":["How", "Are", "You"], "key4": {"name": "xieWei", "age":20,"city":"shanghai"}}`)

	marshalInterface([]byte(name))

}

func boolToJson(ok bool) []byte {

	jsonResult, _ := json.Marshal(ok)
	return jsonResult
}

func jsonToBool(value []byte) bool {
	var goResult bool
	json.Unmarshal(value, &goResult)
	return goResult
}

func intToJson(value int) []byte {
	jsonInt, _ := json.Marshal(value)
	return jsonInt
}

type Info struct {
	Name    string `json:"name,omitempty"`
	Age     int    `json:"age,string"`
	City    string `json:"city_shanghai"`
	Company string `json:"-"`
}

func (i Info) MarshalOp() []byte {
	jsonResult, _ := json.Marshal(i)
	return jsonResult
}

func UnMarshalExample(value []byte) (result Info) {
	json.Unmarshal(value, &result)
	return result

}

type SelfMarshal struct {
	Name string
	Age  int
	City string
}

func (self SelfMarshal) MarshalJSON() ([]byte, error) {
	result := fmt.Sprintf("name:--%s,age:--%d,city:--%s", self.Name, self.Age, self.City)
	if !json.Valid([]byte(result)) {
		fmt.Println("invalid")
		return json.Marshal(result)
	}
	return []byte(result), nil
}

type jsonTime time.Time

//实现它的json序列化方法
func (this jsonTime) MarshalJSON() ([]byte, error) {
	var stamp = fmt.Sprintf("\"%s\"", time.Time(this).Format("2006-01-02 15:04:05"))
	return []byte(stamp), nil
}

type Test struct {
	Date  jsonTime `json:"date"`
	Name  string   `json:"name"`
	State bool     `json:"state"`
}

func marshalInterface(value []byte) {
	var result map[string]interface{}
	json.Unmarshal(value, &result)
	fmt.Println(result)
	for key, value := range result {
		switch value.(type) {
		case float64:
			fmt.Println("Int", result[key], key, reflect.TypeOf(result[key]))
		case string:
			fmt.Println("String", result[key], key, reflect.TypeOf(result[key]))
		case map[string]interface{}:
			fmt.Println("Info", result[key], key, reflect.TypeOf(result[key]))
		case []interface{}:
			fmt.Println("[]string", result[key], key, reflect.TypeOf(result[key]))
		}
	}

}
