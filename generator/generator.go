package generator

import (
	"fmt"
	"os"
	"strconv"
	"text/template"
)

// GenerateCode generates the code for both go constant file and static site.
func GenerateCode() {
	// Check if the constant name already exists. Constant name need to be global unique
	platforms := GetPlatforms("platforms/")

	checkValid(platforms)

	generateGoCode(platforms, "templates/constant.go.tpl", "errcode/constant.go")
	generateSiteCode(platforms, "templates/", "ecms-site/")
}

// generateSiteCode generate site code
func generateSiteCode(platforms []Platform, templateDir, outputDir string) {
	// Generate page contain all errors
	allErrorTmpl, err := template.ParseFiles(templateDir + "all-error.md.tpl")
	if err != nil {
		panic(err)
	}

	f, err := os.Create(outputDir + "content/01 总览/all-error.md") // ignore_security_alert
	if err != nil {
		panic(err)
	}

	err = allErrorTmpl.Execute(f, struct {
		Platforms []Platform
	}{
		Platforms: platforms,
	})

	err = f.Close()
	if err != nil {
		panic(err)
	}

	// Generate single page for each error
	for _, platform := range platforms {
		platformMdPath := fmt.Sprintf("%s/content/02 全部错误码/(%s) %s/",
			outputDir, platform.Code, platform.Name)
		// Ensure the directory exists
		err := os.MkdirAll(platformMdPath, os.ModePerm)
		if err != nil {
			panic(err)
		}

		// Create _index.md for platform
		_indexFile, err := os.Create(platformMdPath + "_index.md") // ignore_security_alert
		if err != nil {
			panic(err)
		}
		err = _indexFile.Close()
		if err != nil {
			panic(err)
		}
		// create page for each error.
		for _, module := range platform.Modules {
			for _, specificError := range module.SpecificErrors {
				singleErrorTmpl, err := template.ParseFiles(templateDir + "single-error.md.tpl")
				if err != nil {
					panic(err)
				}

				errCode := platform.Code + module.Code + specificError.Code
				errName := platform.Prefix + module.Prefix + specificError.Name

				f, err := os.Create(platformMdPath + errCode + ".md") // ignore_security_alert
				if err != nil {
					panic(err)
				}

				err = singleErrorTmpl.Execute(f, struct {
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
				})

				err = f.Close()
				if err != nil {
					panic(err)
				}
			}
		}
	}
}

// generateGoCode generate go code
func generateGoCode(platforms []Platform, templateFile, outputFile string) {
	tmpl, err := template.ParseFiles(templateFile)
	if err != nil {
		panic(err)
	}

	f, err := os.Create(outputFile) // ignore_security_alert
	if err != nil {
		panic(err)
	}

	err = tmpl.Execute(f, platforms)

	err = f.Close()
	if err != nil {
		panic(err)
	}
}

// checkValid Check whether there is conflict before rewrite <constant file>, will panic if there is conflict
func checkValid(platforms []Platform) {
	// Check if the constant name already exists. Constant name need to be global unique
	nameCheck := make(map[string]struct{})
	codeCheck := make(map[int32]struct{})
	for _, platform := range platforms {
		for _, module := range platform.Modules {
			for _, specificError := range module.SpecificErrors {
				constantName := platform.Prefix + module.Prefix + specificError.Name
				constantInt, err := strconv.ParseInt(platform.Code+module.Code+specificError.Code, 10, 32)
				if err != nil {
					panic(err)
				}

				_, exist := nameCheck[constantName]
				if exist {
					panic(fmt.Sprintf("Constant name '%v' already exists", constantName))
				}
				_, exist = codeCheck[int32(constantInt)]
				if exist {
					panic(fmt.Sprintf("Constant code '%v' already exists", constantInt))
				}

				nameCheck[constantName] = struct{}{}
				codeCheck[int32(constantInt)] = struct{}{}
			}
		}
	}
}
