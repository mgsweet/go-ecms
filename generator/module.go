package generator

import (
	"fmt"
	"io/ioutil"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

type Module struct {
	Name           string `yaml:"name"`
	Code           string `yaml:"code"`
	Prefix         string `yaml:"prefix"`
	File           string `yaml:"file"`
	SpecificErrors []SpecificError
}

type Modules struct {
	Modules []Module `yaml:"modules"`
}

func GetModules(dir string) ([]Module, error) {
	yamlFile, err := ioutil.ReadFile(filepath.Join(dir, "config.yaml"))
	if err != nil {
		return nil, err
	}

	modules := Modules{}

	err = yaml.Unmarshal(yamlFile, &modules)
	if err != nil {
		return nil, err
	}

	for i, module := range modules.Modules {
		modules.Modules[i].SpecificErrors, err = GetSpecificErrors(filepath.Join(dir, module.File))
		if err != nil {
			fmt.Printf("Cannot get specific errors for module '%v' skip, %v\n", module.Name, err.Error())
		}
	}

	return modules.Modules, nil
}
