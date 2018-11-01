package main

import (
	"fmt"
	"strings"
)

func stringsClean(value string) string {
	newReplacer := strings.NewReplacer("\n", "\t", " ")
	newValue := newReplacer.Replace(value)
	return strings.TrimSpace(newValue)
}

func stringsContains(value string, subString string) bool {
	return strings.Contains(value, subString)
}

func stringsSplit(value string, splitString string) []string {
	return strings.Split(value, splitString)
}

func stringsCount(value string, subString string) int {
	return strings.Count(value, subString)
}

func stringsLowerOrUpper(value string, toLower bool) string {
	if toLower {
		return strings.ToLower(value)
	}
	return strings.ToUpper(value)
}

func stringsFiled(value string) []string {
	return strings.Fields(value)
}

var globalVar struct {
	Name string
	Age  int
}

func main() {
	fmt.Println(stringsLowerOrUpper("1Iike'", true))
	var charOne rune = 'a'

	fmt.Print(string(charOne), int(charOne))

	fields := stringsFiled("How can i become gopher           ")
	for index, field := range fields {
		fmt.Println(index, field, len(field))
	}
	globalVar.Age = 20
	globalVar.Name = "XieWei"

	var test = []struct {
		in  string
		out string
	}{
		{
			in:  "1",
			out: "2",
		},
	}
	fmt.Println(test)
}
