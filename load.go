package utils

import (
	"encoding/json"
	"io/ioutil"
	"path"

	"gopkg.in/yaml.v2"
)

type callback func(interface{})

var loaders = map[string]func([]byte, interface{}) error{
	".json": LoadConfigFormJsonBytes,
	".yaml": LoadConfigFromYamlBytes,
}

func LoadConfigFormJsonBytes(content []byte, obj interface{}) error {
	return json.Unmarshal(content, obj)
}

func LoadConfigFromYamlBytes(content []byte, obj interface{}) error {
	return yaml.Unmarshal(content, obj)
}

func LoadConfig(file string, v interface{}, callback ...callback) {
	content, err := ioutil.ReadFile(file)

	if err != nil {
		panic("Load Config Error " + err.Error())
	}
	loader, ok := loaders[path.Ext(file)]

	if !ok {
		panic("Unknown File Type：" + path.Ext(file))
	}
	if err = loader(content, v); err == nil {
		for _, _callback := range callback {
			_callback(v)
		}
	}
}
