package generator

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func RegisterPlatform() {
	platformDir := "platforms"
	templateDir := "templates"
	reader := bufio.NewReader(os.Stdin)
	platforms := GetPlatforms(platformDir)
	platformCodeLenLimit := 2

	fmt.Println("------ Already registered Platform List ------")
	for _, platform := range platforms {
		fmt.Printf("%v (code: %v, prefix: %v, dir: %v)\n",
			platform.Name, platform.Code, platform.Prefix, platform.Dir)
	}
	fmt.Println()
	fmt.Println("------ Registering new platform ------")
	fmt.Println("Platform name (e,g Organization 组织中心): ")

	name, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}
	name = strings.Trim(name, "\n")

	fmt.Println("Platform prefix (e,g Org): ")
	prefix, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}
	prefix = strings.Trim(prefix, "\n")

	fmt.Println("Platform dir (e,g organization): ")
	dir, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}
	dir = strings.Trim(dir, "\n")

	ps := Platforms{Platforms: platforms}
	code, err := ps.GetAvailableCode(platformCodeLenLimit)
	if err != nil {
		panic(err)
	}

	newPlatform := Platform{
		Name:   name,
		Code:   code,
		Prefix: prefix,
		Dir:    dir,
	}

	if err := newPlatform.Register(platformDir, filepath.Join(templateDir, "platformConfig.yaml.tpl")); err != nil {
		panic(err)
	}
}
