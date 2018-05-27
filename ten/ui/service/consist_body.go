package service

import (
	"errors"

	"strings"

	"github.com/olebedev/config"
)

var (
	ErrParseYamlFileFail = errors.New("parse yaml file error")
	ErrParseYamlKeyFail  = errors.New("parse file key no exist")
)

func ConsistBody(key string, fileName string) (map[string]interface{}, error) {

	var (
		cfg *config.Config
		err error
	)

	if !judgeYamlFile(fileName) {
		return nil, ErrParseYamlFileFail
	}
	cfg, err = config.ParseYamlFile(fileName)
	if err != nil {
		return nil, ErrParseYamlFileFail
	}

	var (
		cfgParam *config.Config
	)
	cfgParam, err = cfg.Get(key)
	if err != nil {
		return nil, ErrParseYamlKeyFail
	}
	var body = make(map[string]interface{})
	body[key] = cfgParam.Root
	return body, nil

}

func judgeYamlFile(fileName string) bool {
	return strings.HasSuffix(fileName, ".yml")
}
