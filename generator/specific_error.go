package generator

import (
	"fmt"
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

func GetSpecificErrors(file string) ([]SpecificError, error) {
	yamlFile, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}

	specificErrors := SpecificErrors{}

	err = yaml.Unmarshal(yamlFile, &specificErrors)
	if err != nil {
		return nil, err
	}

	return specificErrors.SpecificErrors, nil
}

// Check if the specific error is valid
func (se SpecificError) Check() error {
	if se.Suffix == "" {
		return fmt.Errorf("no suffix is not allow for specific error")
	}

	if err := CheckCode(se.Code, 3); err != nil {
		return fmt.Errorf("specific error '%v' code '%v' is not valid, %v",
			se.Suffix, se.Code, err)
	}
	return nil
}
