package generator

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"text/template"
)

// GenerateCode generates the code for both go constant file and static site.
func GenerateCode() {
	platformDir := "platforms"
	templateDir := "templates"
	outputGoCodeDir := "errcode"
	siteDir := "ecms-site"

	// Check if the constant name already exists. Constant name need to be global unique
	platforms := GetPlatforms(platformDir)

	checkValid(platforms)

	generateGoCode(platforms, filepath.Join(templateDir, "constant.go.tpl"),
		filepath.Join(outputGoCodeDir, "constant.go"))
	generateSiteCode(platforms, templateDir, siteDir)
}

// generateSiteCode generate site code
func generateSiteCode(platforms []Platform, templateDir, siteDir string) error {
	// Generate page contain all errors
	if err := generateErrorSummaryPage(platforms, templateDir, siteDir); err != nil {
		return err
	}

	// Generate single page for each error
	for _, platform := range platforms {
		platformOutputDir := filepath.Join(siteDir, "content", "错误码",
			fmt.Sprintf("(%s) %s", platform.Code, platform.Name))
		if err := ensureDirExist(platformOutputDir); err != nil {
			return err
		}

		// Create _index.md for platform
		_indexFile, err := os.Create(filepath.Join(platformOutputDir, "_index.md")) // ignore_security_alert
		if err != nil {
			return err
		}
		if err = _indexFile.Close(); err != nil {
			return err
		}

		// create page for each error.
		for _, module := range platform.Modules {
			for _, specificError := range module.SpecificErrors {
				singleErrorTmpl, err := template.ParseFiles(filepath.Join(templateDir, "single-error.md.tpl"))
				if err != nil {
					return err
				}

				// The following code is mainly for site content generation.
				errCode := platform.Code + module.Code + specificError.Code
				errName := platform.Prefix + module.Prefix + specificError.Suffix

				f, err := os.Create(filepath.Join(platformOutputDir, errCode+".md")) // ignore_security_alert
				if err != nil {
					return err
				}

				if err = singleErrorTmpl.Execute(f, struct {
					PlatformName string
					PlatformCode string
					ModuleName   string
					ModuleCode   string
					ErrorCode    string
					ErrorName    string
					ErrorDesc    string
				}{
					PlatformName: platform.Name,
					PlatformCode: platform.Code,
					ModuleName:   module.Name,
					ModuleCode:   module.Code,
					ErrorCode:    errCode,
					ErrorName:    errName,
					ErrorDesc:    specificError.Desc,
				}); err != nil {
					return err
				}

				if err = f.Close(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

// generateErrorSummaryPage Generate a page contain all errors
func generateErrorSummaryPage(platforms []Platform, templateDir, siteDir string) error {
	allErrorTmpl, err := template.ParseFiles(filepath.Join(templateDir, "all-error.md.tpl"))
	if err != nil {
		return err
	}

	outputDir := filepath.Join(siteDir, "content", "总览")
	if err := ensureDirExist(outputDir); err != nil {
		return err
	}

	f, err := os.Create(filepath.Join(outputDir, "all-error.md")) // ignore_security_alert
	if err != nil {
		return err
	}

	if err = allErrorTmpl.Execute(f, struct {
		Platforms []Platform
	}{
		Platforms: platforms,
	}); err != nil {
		return err
	}

	if err = f.Close(); err != nil {
		return err
	}

	return nil
}

func generateGoCode(platforms []Platform, templateFile, outputFile string) error {
	tmpl, err := template.ParseFiles(templateFile)
	if err != nil {
		return err
	}

	f, err := os.Create(outputFile) // ignore_security_alert
	if err != nil {
		return err
	}

	if err = tmpl.Execute(f, platforms); err != nil {
		return err
	}

	if err = f.Close(); err != nil {
		return err
	}

	return nil
}

func ensureDirExist(dir string) error {
	err := os.MkdirAll(dir, os.ModePerm)
	if err != nil {
		return err
	}
	return nil
}

// checkValid Check whether there is conflict before rewrite <constant file>, will panic if there is conflict
func checkValid(platforms []Platform) error {
	// Check if the constant name already exists. Constant name need to be global unique
	nameCheck := make(map[string]struct{})
	codeCheck := make(map[int32]struct{})
	for _, platform := range platforms {
		for _, module := range platform.Modules {
			for _, specificError := range module.SpecificErrors {
				constantName := platform.Prefix + module.Prefix + specificError.Suffix
				constantCode, err := strconv.ParseInt(platform.Code+module.Code+specificError.Code, 10, 32)
				if err != nil {
					return fmt.Errorf("the composed error code of %v cannot be parsed into int32. (platform: %v, module: %v, error suffix: %v)",
						constantName, platform.Name, module.Name, specificError.Suffix)
				}

				_, exist := nameCheck[constantName]
				if exist {
					return fmt.Errorf("constant name '%v' already exists", constantName)
				}
				_, exist = codeCheck[int32(constantCode)]
				if exist {
					return fmt.Errorf("constant code '%v' already exists", constantCode)
				}

				nameCheck[constantName] = struct{}{}
				codeCheck[int32(constantCode)] = struct{}{}
			}
		}
	}
	return nil
}
