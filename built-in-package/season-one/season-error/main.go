package main

import (
	"errors"
	"fmt"
	"net/http"
	"reflect"
	"strconv"
)

func main() {

	recordError := errors.New("record not found")
	dbError := fmt.Errorf("%s", "db connet fail")
	fmt.Println(recordError, reflect.TypeOf(recordError))
	fmt.Println(dbError, reflect.TypeOf(dbError))

	var codeError CodeError
	var err error
	codeError = CodeError{
		Code:    404,
		Message: "http status code error",
	}
	err = codeError
	fmt.Println(err, reflect.TypeOf(err), reflect.TypeOf(err).Kind())

	number, err := strconv.Atoi("2992-121")
	fmt.Println(number, err)

	response, err := http.Get("https://space.bilibili.com/10056291/#/")
	fmt.Println(response, err)
}

type CodeError struct {
	Code    int
	Message string
}

func (c CodeError) Error() string {
	return fmt.Sprintf(" e.Code = %d, e.Message=%s", c.Code, c.Message)
}

func Result() error {
	var codeError CodeError
	codeError = CodeError{
		Code:    400,
		Message: "connect fail",
	}
	return codeError
}

type ExampleError struct {
	Err     error
	Code    int
	Message string
}

func (e *ExampleError) Error() string {
	return fmt.Sprintf("e.Code = %d e.Err = %s e.Message = %s", e.Code, e.Err.Error(), e.Message)
}

func ExampleResult() error {
	return &ExampleError{
		Err:     errors.New("what the fucking life"),
		Code:    502,
		Message: "what the fucking life",
	}
}
