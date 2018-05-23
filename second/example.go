package main

import (
	"fmt"
	"math"
	"strings"
)

/*

1. 变量
*/

func numberExample() {

	var (
		number         int
		numberOfClass  int
		numberOfSchool int
		Number         int
		ok             bool
		price          float64
	)

	var (
		ListOfPeopleName = []string{"xiewei", "xiewei2", "xiewei3"}
		LastIndex        = len(ListOfPeopleName) - 1
		FirstIndex       = 0
	)
	/* bool
		ok
		found
		success
	    flag
	    done

	*/
	fmt.Println(number, numberOfClass, numberOfSchool, Number, ok, price,
		ListOfPeopleName[FirstIndex], ListOfPeopleName[LastIndex])

	var value string
	value = "Hello Golang"
	fmt.Println(value)

	//valueTwo := "xiewei"
	//fmt.Println(valueTwo)

}

/*
2. const

*/

const Name = "xieWei is a man"

func PrintName() {
	fmt.Println(Name)
}

/*
3. string

*/

func ExampleString() {

	var (
		ShortString string
		LongString  string
	)
	ShortString = "xieWei is learning golang"
	LongString = `
Hello Everyone:
  Welcome to here.
  Today, we will learn how to learn golang.
`
	fmt.Println(ShortString, LongString)
	fmt.Println(strings.Contains(ShortString, "xen")) //false
	fmt.Println(strings.Contains(ShortString, "xie")) //true

	fmt.Println(strings.Count(ShortString, "x")) //1

	fmt.Println(strings.LastIndex(ShortString, "g")) //24

	fmt.Println(strings.ToUpper(ShortString)) // XIEWEI IS LEARNING GOLANG
	fmt.Println(strings.Split(ShortString, "g"))
	fmt.Println(strings.Join([]string{"1", "2", "3", "4"}, "")) //1234
	fmt.Println(strings.Replace(ShortString, "xie", "zen", -1)) // zenWei is learning golang

	newReplacer := strings.NewReplacer("x", "X", "i", "I", "w", "W", "e", "E")
	var newShortString string
	newShortString = newReplacer.Replace(ShortString)
	fmt.Println(newShortString) // XIEWEI is learning golang

}

/*
4. List


*/

func ExampleList() {

	var NewListInt = [4]int{1, 2, 3, 4}
	var NewListString = [3]string{"1", "2", "3"}

	fmt.Println(NewListInt, NewListString, NewListInt[0], NewListString[1], len(NewListInt), cap(NewListString))

	for index, value := range NewListInt {
		fmt.Println(index, value+2)
	}

}

/*

5. slice
*/

func ExampleSlice() {
	var (
		NewSliceInt = []int{1, 2}
	)
	NewSliceInt = append(NewSliceInt, 100)
	fmt.Println(NewSliceInt, len(NewSliceInt), NewSliceInt[2]) // [1, 2, 100], 3, 100
}

/*

6. map

*/

func ExampleMap() {

	var (
		NewMap = map[string]int{}
	)
	NewMap["1"] = 2
	NewMap["key"] = 100
	fmt.Println(NewMap) // map[1:2, key:100]

	for key, value := range NewMap {
		fmt.Println(key, value+1000)
	}

}

/*
7. struct
*/

type Info struct {
	name   string
	age    int
	school string
	habits map[string]string
}

func ExampleStruct() {

	var (
		newInfo Info
		habit   map[string]string
	)
	habit = make(map[string]string)
	newInfo.name = "xiewei"
	newInfo.age = 18
	newInfo.school = "xxx"
	habit["xxx"] = "xxx"
	habit["yyy"] = "yyy"
	newInfo.habits = habit
	fmt.Println(newInfo, newInfo.name, newInfo.GetAttr(), newInfo.SetAttr())

}

func (i Info) GetAttr() interface{} {
	return i.name
}

func (i *Info) SetAttr() interface{} {
	i.name = "XieWei"
	return i.name
}

/*

8. func

*/

func ExampleFunc(argOne float64, argTwo float64, name string, value interface{}) float64 {
	return math.Max(argOne, argTwo)
}

func ExampleFunc1(argOne ...float64) float64 {
	return 0
}

func ExampleNoName() {

	var (
		PrintName = func(name string) string { return name }
	)

	fmt.Println(PrintName("xiweiwiejweiw"))
}

/*
9. if...else
   switch
   for


*/

func ExampleJudgement(name string) bool {
	if name == "" {
		return false
	} else {
		return true
	}

}

func ExampleFor(name []string) {

	if name == nil {
		return
	}
	for i := 0; i < len(name); i++ {
		fmt.Println(name[i])
	}
}

func main() {
	numberExample()
	PrintName()
	ExampleString()
	ExampleList()
	ExampleSlice()
	ExampleMap()
	fmt.Println("________________________")
	ExampleStruct()

	fmt.Println(ExampleFunc(1, 2))

	ExampleNoName()

	fmt.Println(ExampleJudgement("xiewei"))
	ExampleFor([]string{"xiewei", "wenlimin", "xiewei2", "wenlimin2"})

}
