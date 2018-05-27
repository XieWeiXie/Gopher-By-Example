package roles

import (
	"errors"
	"fmt"
	"go-example-for-live/ten/domain/constants"
	"go-example-for-live/ten/infra/golang/osadapater"
	"go-example-for-live/ten/ui/service"
)

var (
	ErrInstanceIdExist  = errors.New("config instance exist")
	ErrInstaceeMakeFail = errors.New("make instance fail")
)

func CreateConfigDir(instanceId string, fullPath string) (bool, error) {

	var path string
	path = constants.ProjectPath + instanceId
	//fmt.Println(os.Getwd())
	//fmt.Println(path)
	var found bool
	found = osadapater.DExist(path)
	if found != false {
		return false, ErrInstanceIdExist
	}

	var err error

	err = osadapater.MKdir(path) // wrong
	fmt.Println(err)
	if err != nil {
		return false, ErrInstaceeMakeFail
	}
	return service.GitClone(path, "github.com", fullPath)

}

func RemoveConfigDir(instanceId string) (bool, error) {

	var path string
	path = constants.ProjectPath + instanceId

	var err error
	err = osadapater.RDir(path)
	if err != nil {
		return false, err
	}
	return true, nil
}
