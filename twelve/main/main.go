package main

import "fmt"

func main() {

	// map、slice、channel

	var number int

	number = 0
	copyNumber := number
	copyNumber = 3

	fmt.Printf("----: new=%v, old=%v\n", copyNumber, number)

	var List = [2]int{1, 2}

	copyList := List
	copyList[1] = 4

	fmt.Printf("----: new=%v, old=%v\n", copyList, List)

	type Info struct {
		Name   string
		Age    int
		School string
	}

	oldInfo := Info{
		"xieWei", 20, "shanghaiUniversity",
	}

	newInfo := oldInfo
	newInfo.Name = "xieWei2"

	fmt.Printf("----: new=%v, old=%v\n", newInfo, oldInfo)

	var Body = make(map[string]string)
	Body["name"] = "xieWei"
	Body["School"] = "shanghaiUniversity"

	newBody := Body
	newBody["name"] = "xieWei2"
	newBody["Age"] = "20"

	fmt.Printf("----: new=%v, old=%v\n", newBody, Body)

	var oldSlice = []int{1, 2, 3, 4}

	newSlice := oldSlice
	newSlice[3] = 10
	newSlice = append(newSlice, 100)

	fmt.Printf("----: new=%v, old=%v\n", newSlice, oldSlice)

	var exampleSlice = make([]int, 4)
	exampleSlice[0] = 1
	fmt.Println(exampleSlice)

	var exampleInt int

	exampleInt = 4

	newExampleInt := &exampleInt
	*newExampleInt = 40

	fmt.Println(exampleInt, *newExampleInt)

	exampleBigInfo := new(Info)
	fmt.Println(exampleBigInfo)

	fmt.Println(exampleBigInfo.Age)
}
