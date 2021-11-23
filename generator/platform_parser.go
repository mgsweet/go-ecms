package generator

import (
	"io/ioutil"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

type Platform struct {
	Name    string `yaml:"name"`
	Code    string `yaml:"code"`
	Prefix  string `yaml:"prefix"`
	Dir     string `yaml:"dir"`
	Modules []Module
}

type Platforms struct {
	Platforms []Platform `yaml:"platforms"`
}

func GetPlatforms(dir string) []Platform {
	yamlFile, err := ioutil.ReadFile(filepath.Join(dir, "config.yaml"))
	if err != nil {
		panic(err)
	}

	platforms := Platforms{}

	err = yaml.Unmarshal(yamlFile, &platforms)
	if err != nil {
		panic(err)
	}

	for i, platform := range platforms.Platforms {
		platforms.Platforms[i].Modules = GetModules(filepath.Join(dir, platform.Dir))
	}

	return platforms.Platforms
}
