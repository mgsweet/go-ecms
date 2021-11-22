package generator

import (
	"io/ioutil"

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

func GetModules(dir string) []Module {
	yamlFile, err := ioutil.ReadFile(dir + "config.yaml")
	if err != nil {
		panic(err)
	}

	modules := Modules{}

	err = yaml.Unmarshal(yamlFile, &modules)
	if err != nil {
		panic(err)
	}

	for i, module := range modules.Modules {
		modules.Modules[i].SpecificErrors = GetSpecificErrors(dir + module.File)
	}

	return modules.Modules
}
