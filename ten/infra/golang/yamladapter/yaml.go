package yamladapter

import "gopkg.in/yaml.v2"

func Marshal(v interface{}) (out []byte, err error) {
	out, err = yaml.Marshal(v)
	return
}

func Unmarshal(in []byte, out interface{}) error {
	return yaml.Unmarshal(in, out)
}
