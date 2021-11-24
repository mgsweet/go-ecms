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

// GetModules returns a list of modules by parsing files in the given directory
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

// Check checks if the module is valid
func (m *Module) Check() error {
	if m.Prefix == "" {
		return fmt.Errorf("no prefix is not allow for module: %v", m.Name)
	}

	if err := CheckCode(m.Code, 3); err != nil {
		return fmt.Errorf("module '%v' code '%v' is not valid, %v", m.Name, m.Code, err)
	}
	return nil
}
