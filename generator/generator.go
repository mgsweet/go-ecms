package generator

import (
	"fmt"
	"os"
	"strconv"
	"text/template"
)

func GenerateGoCode(inputDir, templateFile, outputFile string) {
	// Check if the constant name already exists. Constant name need to be global unique
	nameCheck := make(map[string]struct{})
	codeCheck := make(map[int32]struct{})
	platforms := GetPlatforms(inputDir)
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

	tmpl, err := template.ParseFiles(templateFile)
	if err != nil {
		panic(err)
	}

	f, err := os.Create(outputFile)
	if err != nil {
		panic(err)
	}

	err = tmpl.Execute(f, platforms)
}
