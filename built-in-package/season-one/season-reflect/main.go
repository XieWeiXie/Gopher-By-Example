package main

import (
	"fmt"
	"reflect"
	"strings"
)

type Info struct {
	Name  string      `json:"name"`
	Age   interface{} `json:"age"`
	Price float32     `json:"price"`
}

type Groups struct {
	ID        uint     `json:"id"`
	Name      string   `json:"name"`
	Count     int      `json:"count"`
	CanFly    bool     `json:"can_fly"`
	GroupIDs  []int    `json:"group_ids"`
	GroupName []string `json:"group_name"`
	Info      `json:"info"`
}

var globalValue struct {
	groups Groups
}

func init() {
	globalValue.groups = Groups{
		ID:        1,
		Name:      "xieWei",
		Count:     12,
		CanFly:    false,
		GroupIDs:  []int{100, 200, 300, 400},
		GroupName: []string{"what", "how", "when", "why"},
		Info: Info{
			Name:  "XieXiaoLu",
			Age:   20,
			Price: 1.2345,
		},
	}

	valid.tag = "xieWei"

	valid.value = Model{
		ID:    1,
		Count: 2000,
		Name:  "Golang",
	}

}
func (g Groups) TypeHandler() {
	typeGroups := reflect.TypeOf(g)
	name := typeGroups.Name()
	fmt.Println("Name: ", name, "Kind", typeGroups.Kind())
	for i := 0; i < typeGroups.NumField(); i++ {
		filed := typeGroups.Field(i)
		fmt.Println(filed.Name, "\t", filed.Tag, "\t", reflect.ValueOf(filed), "\t", filed.Type)
	}

	for i := 0; i < typeGroups.NumField(); i++ {
		filedByIndex := typeGroups.FieldByIndex([]int{i})
		filedByName, _ := typeGroups.FieldByName(filedByIndex.Name)
		fmt.Println(filedByIndex, filedByIndex.Name, filedByIndex.Type)
		fmt.Println(filedByName, filedByName.Name, filedByName.Type)
	}

	for i := 0; i < typeGroups.NumMethod(); i++ {
		method := typeGroups.Method(i)
		fmt.Println(method.Name, method.Type)
	}
}

func (g Groups) ValueHandler() {
	valueGroup := reflect.ValueOf(g)
	fmt.Println(valueGroup.NumField(), valueGroup.NumMethod(), valueGroup, reflect.ValueOf(&g).Elem())

	for i := 0; i < valueGroup.NumField(); i++ {
		field := valueGroup.Field(i)
		fmt.Println(field, field.Type(), field.Kind())
	}

	method := valueGroup.MethodByName("TypeHandler")
	fmt.Println(method, method.Kind(), method.Type())

	for i := 0; i < valueGroup.NumMethod(); i++ {
		method := valueGroup.Method(i)
		fmt.Println(method.Type())
	}
	ref := reflect.ValueOf(&g).Elem()

	fmt.Println(ref.FieldByName("Name"), ref.Field(0))
}

var valid struct {
	tag   string
	value Model
}

type Model struct {
	ID     uint   `xieWei:"number,max=10,min=1"`
	Name   string `xieWei:"string"`
	Count  int    `xieWei:"number,max=100,min=1"`
	CanFly bool   `xieWei:"bool,default=false"`
}

func (m Model) Handler(name string) bool {

	typeModel := reflect.TypeOf(m)
	if tag, ok := typeModel.FieldByName(name); ok {
		if ok := strings.HasPrefix(string(tag.Tag), valid.tag); ok {
			//fmt.Println(validTagList[0])
			validTagList := strings.FieldsFunc(string(tag.Tag), func(r rune) bool {
				return r == ',' || r == '"'
			})
			switch validTagList[1] {
			case "number":
				{
					fmt.Println(validTagList[1:])
				}
			case "string":
				fmt.Println(validTagList[1:])

			case "bool":
				fmt.Println(validTagList[1:])

			}

		} else {
			return false
		}
	}
	return false
}

func main() {
	//globalValue.groups.TypeHandler()
	//globalValue.groups.ValueHandler()
	valid.value.Handler("ID")
	valid.value.Handler("Name")
	valid.value.Handler("Count")
	valid.value.Handler("CanFly")
	fmt.Println(stringHandler(`validate:"number,min=1,max=1000"`)[1:])

	var number int
	number = 1
	fmt.Println(reflect.TypeOf(number), reflect.ValueOf(number))

	type numberExample int
	var numberOther numberExample
	numberOther = 2
	fmt.Println(reflect.TypeOf(numberOther), reflect.ValueOf(numberOther), reflect.ValueOf(numberOther).Kind())
}

func stringHandler(value string) []string {
	return strings.FieldsFunc(value, func(r rune) bool {
		return r == '"' || r == ','
	})
}

// TODO

/*
1. 结构体获取类型、获取 tag, 获取属性的 tag, 有效性是如何操作？获取 tag, 得到字符串
2.

*/
