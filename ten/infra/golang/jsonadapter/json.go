package jsonadapter

import (
	"encoding/json"
	"errors"
	"fmt"
)

// 功能点： json.Marshal and json.Unmarshal

var (
	ErrUnmarshal = errors.New("json unmarshal fail")
	ErrMarshal   = errors.New("json marshal fail")
)

func Unmarshal(data []byte, v interface{}) error {

	var err error
	err = json.Unmarshal(data, &v)
	if err != nil {
		fmt.Println(ErrUnmarshal)
		return ErrUnmarshal
	}
	return nil

}

func Marshal(v interface{}) ([]byte, error) {

	var (
		res []byte
		err error
	)

	res, err = json.Marshal(v)
	if err != nil {
		fmt.Println(ErrMarshal)
		return nil, ErrMarshal
	}
	return res, nil
}
