package app

import (
	"errors"
	"fmt"
	"go-example-for-live/ten/domain/roles"
	"go-example-for-live/ten/infra/golang/ioutiladapter"
	"go-example-for-live/ten/infra/golang/yamladapter"
	"go-example-for-live/ten/ui/service"
	"strings"
)

var (
	InstanceError = errors.New("instance error Or already exist")
)

type ConfigInstance struct {
	InstanceName string
}

func NewConfigInstance(name string) *ConfigInstance {
	return &ConfigInstance{
		InstanceName: name,
	}
}

func (c *ConfigInstance) Instance(fullPath string) {
	// create config instance
	var (
		ok  bool
		err error
	)

	ok, err = roles.CreateConfigDir(c.InstanceName, fullPath)
	if err != nil {
		fmt.Println(InstanceError, ok)
	}

	// consist boy
	var (
		structTreeConfig map[string]interface{}
	)
	structTreeConfig = service.GKeyAndFile("example")
	cfgFileName := "go-example-for-live/ten/config/cfg.yml"
	for key, value := range structTreeConfig {
		var (
			newKey   string
			newValue string
		)
		newKey = splitKey(key)
		newValue = splitValue(value.(string))
		//fmt.Println(newKey, newValue)

		var body = make(map[string]interface{})
		var err error
		body["parameters"], err = service.ConsistBody(newKey, cfgFileName)
		if err != nil {

		}
		fmt.Println(body)

		out, _ := yamladapter.Marshal(body)
		ioutiladapter.WriteCfgIntoFile("go-example-for-live/ten/ConfigResume/"+c.InstanceName+"/"+
			"Resume/"+newValue+".yml", out)
	}

}

func splitKey(key string) string {
	var (
		stringList []string
	)
	stringList = strings.Split(key, ".")
	return stringList[0]
}

func splitValue(key string) string {
	var (
		stringList []string
	)

	stringList = strings.Split(key, ".")
	if len(stringList) <= 2 {
		return stringList[1]
	} else {
		return stringList[2]
	}

}
