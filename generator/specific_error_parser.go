package generator

import (
	"io/ioutil"

	"gopkg.in/yaml.v3"
)

type SpecificError struct {
	Suffix string `yaml:"suffix"`
	Code   string `yaml:"code"`
	Desc   string `yaml:"desc"`
}

type SpecificErrors struct {
	SpecificErrors []SpecificError `yaml:"specific_errors"`
}

func GetSpecificErrors(file string) []SpecificError {
	yamlFile, err := ioutil.ReadFile(file)
	if err != nil {
		panic(err)
	}

	specificErrors := SpecificErrors{}

	err = yaml.Unmarshal(yamlFile, &specificErrors)
	if err != nil {
		panic(err)
	}

	return specificErrors.SpecificErrors
}
