package utils

import (
	"os"

	"gopkg.in/yaml.v2"
)

func InitConfig(filePath string, configObjectRef interface{}) error {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return err
	}

	err = yaml.Unmarshal(data, configObjectRef)
	if err != nil {
		return err
	}

	return nil
}
