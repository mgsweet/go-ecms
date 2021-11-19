package main

import (
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v3"
)

type Platform struct {
	name    string `yaml:"name"`
	code    string `yaml:"code"`
	prefix  string `yaml:"prefix"`
	dirPath string `yaml:"dir_path"`
}

type Platforms struct {
	Platforms []Platform `yaml:"platforms"`
}

func main() {
	yPlatformConfig, err := ioutil.ReadFile("platforms/config.yaml")
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	platforms := Platforms{}

	err = yaml.Unmarshal(yPlatformConfig, &platforms)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	for _, platform := range platforms.Platforms {
		fmt.Println(platform)
	}
}
