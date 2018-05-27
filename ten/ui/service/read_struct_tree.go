package service

import (
	"errors"
	"fmt"

	"go-example-for-live/ten/domain/constants"

	"github.com/olebedev/config"
)

var (
	ErrConfigKeyFail = errors.New("config key fail")
)

func GKeyAndFile(key string) map[string]interface{} {

	var (
		cfg *config.Config
		err error
	)

	cfg, err = config.ParseYamlFile(constants.StructTreePath)

	var (
		rt map[string]interface{}
	)
	rt, err = cfg.Map("parameters." + key)
	if err != nil {
		fmt.Println(ErrConfigKeyFail)
		return nil
	}
	return rt
}
