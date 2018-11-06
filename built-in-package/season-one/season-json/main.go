package main

import (
	"encoding/json"
	"fmt"
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
		Name: "XieWei",
		Age:  100,
		City: "shangHai",
	}
	fmt.Println(string(info.MarshalOp()))

	var otherInfo Info
	otherInfo.Name = ""
	otherInfo.Age = 20
	otherInfo.City = "BeiJing"
	fmt.Println(string(otherInfo.MarshalOp()))

	fmt.Println(UnMarshalExample([]byte(`{"name":"xiewei", "age": "20", "city_shanghai": "Guangdong"}`)))

	var self = SelfMarshal{}
	self.Age = 20
	self.Name = "XieWei"
	self.City = "HangZhou"
	fmt.Println(self, 12)

	selfjsjjadjsajdjaj, err := json.Marshal(self)
	fmt.Println(err, string(selfjsjjadjsajdjaj))

	var t = Test{}
	t.Date = jsonTime(time.Now())
	t.Name = "Hello World"
	t.State = true
	body, _ := json.Marshal(t)
	fmt.Println(string(body))

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
	Name string `json:"name,omitempty"`
	Age  int    `json:"age,string"`
	City string `json:"city_shanghai"`
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
	result := fmt.Sprintf("name:%s,age:%d,city:%s", self.Name, self.Age, self.City)
	fmt.Println(result, 98765432)
	return json.Marshal(result)
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
