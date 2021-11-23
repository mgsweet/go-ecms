package generator

import (
	"fmt"
	"os"
	"path/filepath"
	"text/template"
)

// GenerateCode generates the code for both go constant file and static site based on given yaml configurations.
func GenerateCode() {
	platformDir := "platforms"
	templateDir := "templates"
	outputGoCodeDir := "errcode"
	siteDir := "ecms-site"

	// parse yaml files
	platforms := GetPlatforms(platformDir)

	// do some validation
	if err := CheckValid(platforms); err != nil {
		panic(err)
	}

	// generate go code
	if err := generateGoCode(platforms, filepath.Join(templateDir, "constant.go.tpl"),
		filepath.Join(outputGoCodeDir, "constant.go")); err != nil {
		panic(err)
	}

	// generate site code
	if err := generateSiteCode(platforms, templateDir, siteDir); err != nil {
		panic(err)
	}
}

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
