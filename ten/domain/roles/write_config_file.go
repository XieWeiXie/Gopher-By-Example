package roles

import (
	"errors"
	"fmt"
	"go-example-for-live/ten/infra/golang/ioutiladapter"
	"go-example-for-live/ten/infra/golang/yamladapter"
)

var (
	ErrYamlParseFail   = errors.New("yaml adapter fail")
	ErrWriteConfigFail = errors.New("ioutil adapter fail")
)

func WriteConfigIntoFile(data map[string]interface{}, file string) (bool, error) {

	var (
		newData []byte
		err     error
	)

	newData, err = yamladapter.Marshal(data)
	if err != nil {
		fmt.Println(ErrYamlParseFail)
		return false, ErrYamlParseFail
	}

	err = ioutiladapter.WriteCfgIntoFile(file, newData)
	if err != nil {
		fmt.Println(ErrWriteConfigFail)
		return false, ErrWriteConfigFail
	}
	return true, nil
}
