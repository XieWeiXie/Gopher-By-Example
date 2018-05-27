package ioutiladapter

import "io/ioutil"

func WriteCfgIntoFile(fileName string, data []byte) error {

	return ioutil.WriteFile(fileName, data, 0644)
}

func ReadAllCfgFromFile(fileName string) ([]byte, error) {

	var (
		out []byte
		err error
	)
	out, err = ioutil.ReadFile(fileName)
	return out, err
}
