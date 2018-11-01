package main

import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/pkg/errors"
)

var globalIntVal struct {
	i int
	j int8
	k int16
	l int32
	g int64
}

var globalUintVal struct {
	i uint
	j uint8
	k uint16
	l uint32
	g uint64
}

var globalFloatVal struct {
	i float32
	j float64
}

var globalRuneVal struct {
	i rune
}

type StrError struct {
	Func string
	Num  string
	Err  error
}

func (s StrError) Error() string {
	return fmt.Sprintf("func: %s, num: %s, Err: %s", s.Func, s.Num, s.Err.Error())
}

func example() {
	strconv.Itoa(1)
	strconv.Atoi("1")
	q := strconv.QuoteToASCII("Hello, 世界")
	fmt.Println(q)
	rst := make([]byte, 0)
	rst = strconv.AppendBool(rst, 0 < 1)
	fmt.Printf("%s%d\n", rst, len(rst)) // true
	rst = strconv.AppendBool(rst, 0 > 1)
	fmt.Printf("%s%d\n", rst, len(rst)) // truefalse

	var e StrError
	e.Func = "123"
	e.Num = "456"
	e.Err = errors.New("dadada")

	var ee error
	ee = e

	fmt.Println(ee.(StrError))

	strconv.ParseFloat("21212", 10)

}

func toInt(value string) (result int) {
	result, _ = strconv.Atoi(value)
	return

}

func toString(value int) (result string) {
	result = strconv.Itoa(value)
	return
}

func toBool(value string) (result bool) {
	result, _ = strconv.ParseBool(value)
	return
}

func boolToString(value bool) (result string) {
	result = strconv.FormatBool(value)
	return
}

func toFloat(value string) (result float64) {
	result, _ = strconv.ParseFloat(value, 32)
	return
}

func floatToString(value float64) (result string) {
	result = strconv.FormatFloat(value, 'b', -1, 32)
	return
}

func toBaseInt(value string) (result int64) {
	result, _ = strconv.ParseInt("123", 8, 32)
	return
}

func IntToString(value int64) (result string) {
	return strconv.FormatInt(value, 8)
}

func main() {
	fmt.Println(toInt("123"), reflect.TypeOf(toInt("123")))
	fmt.Println(toString(1233333), reflect.TypeOf(toString(1233333)))

	fmt.Println(toBool("True"), reflect.TypeOf(toBool("TRUE")))
	fmt.Println(boolToString(true), reflect.TypeOf(boolToString(false)))

	fmt.Println(toFloat("123.333"), reflect.TypeOf(toFloat("12")))
	fmt.Println(floatToString(12.23444), reflect.TypeOf(floatToString(12.22)))

	fmt.Println(toBaseInt("123"))
	fmt.Println(IntToString(123))

	fmt.Println(strconv.Quote("1234444"), strconv.QuoteRune('b'))
	fmt.Println(strconv.Quote("1234444"), strconv.QuoteRune('b'))
}
