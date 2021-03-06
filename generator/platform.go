package generator

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
	"text/template"

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

// GetAvailableCode returns an available code for a new platform.
// Currently, it will only take the current maximum value plus one
func (ps *Platforms) GetAvailableCode(length int) (string, error) {
	platforms := ps.Platforms

	// non-common platforms start from 20
	var nextCodeInt int64 = 20
	for _, platform := range platforms {
		codeInt, err := strconv.ParseInt(platform.Code, 10, 32)
		if err != nil {
			return "", err
		}
		if codeInt >= nextCodeInt {
			nextCodeInt = codeInt + 1
		}
	}

	nextCode := strconv.FormatInt(nextCodeInt, 10)
	if len(nextCode) != length {
		return "", fmt.Errorf("adding a number after the largest code exceeds the length limit: %d", length)
	}

	return nextCode, nil
}

// GetPlatforms returns a list of platforms by parsing files in the given directory
func GetPlatforms(dir string) ([]Platform, error) {
	yamlFile, err := ioutil.ReadFile(filepath.Join(dir, "config.yaml"))
	if err != nil {
		return nil, err
	}

	platforms := Platforms{}

	err = yaml.Unmarshal(yamlFile, &platforms)
	if err != nil {
		return nil, err
	}

	for i, platform := range platforms.Platforms {
		platforms.Platforms[i].Modules, err = GetModules(filepath.Join(dir, platform.Dir))
		if err != nil {
			fmt.Printf("Cannot get module from platform '%v' skip, %v\n", platform.Name, err.Error())
		}
	}

	return platforms.Platforms, nil
}

// SavePlatformsToYaml saves the platforms to a yaml file
func (ps *Platforms) SavePlatformsToYaml(templateFile, outputFile string) error {
	tmpl, err := template.ParseFiles(templateFile)
	if err != nil {
		return err
	}

	f, err := os.Create(outputFile) // ignore_security_alert
	if err != nil {
		return err
	}

	if err = tmpl.Execute(f, ps.Platforms); err != nil {
		return err
	}

	if err = f.Close(); err != nil {
		return err
	}

	return nil
}

// Register create platform dir and edit config.yaml in the given directory
func (p *Platform) Register(platformDir, configTemplateFile string) error {

	// Validate the platform
	if err := p.Check(); err != nil {
		return err
	}
	// Check if the platform already exists
	platforms, err := GetPlatforms(platformDir)
	if err != nil {
		return err
	}
	for _, platform := range platforms {
		if platform.Name == p.Name {
			return fmt.Errorf("Platform already registered, cannot register duplicate platform")
		}
	}

	platforms = append(platforms, *p)

	// Create platform directory
	if err := EnsureDirExist(filepath.Join(platformDir, p.Dir)); err != nil {
		return err
	}

	// Edit config.yaml
	ps := Platforms{
		Platforms: platforms,
	}
	if err := ps.SavePlatformsToYaml(configTemplateFile, filepath.Join(platformDir, "config.yaml")); err != nil {
		return err
	}
	return nil
}

// Check if the platform is valid
func (p *Platform) Check() error {
	if p.Code != "10" && p.Prefix == "" {
		return fmt.Errorf("no prefix is not allow for business related platform: %v", p.Name)
	}

	if err := CheckCode(p.Code, 2); err != nil {
		return fmt.Errorf("platform '%v' code '%v' is not valid, %v", p.Name, p.Code, err)
	}

	num, err := strconv.ParseInt(p.Code, 10, 32)
	if err != nil {
		return fmt.Errorf("platform '%v' code '%v' is not valid, cannot parse code to int32",
			p.Name, p.Code)
	}

	if num != 10 && num < 20 {
		return fmt.Errorf("platform '%v' code '%v' is not valid, business related platform code can not smaller than 20",
			p.Name, p.Code)
	}
	return nil
}
