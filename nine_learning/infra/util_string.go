package infra

import (
	"regexp"
	"strconv"
	"strings"
)

func StringTrim(old string) string {
	return strings.TrimSpace(old)
}

func StringSplit(old string) string {

	if old == "None" {
		return "None"
	}

	var (
		stringList []string
		length     int
	)
	stringList = strings.Split(old, "/")
	length = len(stringList)

	return stringList[length-1]

}

func StringsReplace(old string) string {

	var NewReplacer *strings.Replacer

	NewReplacer = strings.NewReplacer("\t", "")
	return NewReplacer.Replace(old)
}

func StringSplitByDot(old string) (int, int, int) {
	var stringList []string

	stringList = strings.Split(old, "•")

	var numberList []int
	for index, oneString := range stringList {
		if index > 0 && index < 4 {
			number := StringRegexp(oneString)
			numberList = append(numberList, number)
			//fmt.Println(number)
		}
	}
	//fmt.Println(numberList)
	return numberList[0], numberList[1], numberList[2]
	//return 0, 0, 0
}

func StringRegexp(old string) int {
	reg := regexp.MustCompile(`\d+`)
	stringContent := reg.FindString(old)
	number, _ := strconv.Atoi(stringContent)
	return number
}
