package common

import (
	"io/ioutil"

	yaml "gopkg.in/yaml.v1"
)

func ReadConf(filepath string, out *interface{}) error {
	conffile, err := ioutil.ReadFile(filepath)
	if err != nil {
		return err
	}

	return yaml.Unmarshal(conffile, &out)
}
