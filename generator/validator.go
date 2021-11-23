package generator

import (
	"fmt"
	"strconv"
)

// CheckValid Check whether there is conflict before rewrite <constant file>, will panic if there is conflict
// Rules:
// 1. Constant error name need to be global unique
// 2. Constant error code need to be global unique
// 3. No prefix is not allow for business related platform
// 4. BB CCC DDD code format should be followed
// 5. No prefix is not allow for module
func CheckValid(platforms []Platform) error {
	nameUniqueCheck := make(map[string]struct{})
	codeUniqueCheck := make(map[string]struct{})
	platformPrefixUniqueCheck := make(map[string]struct{})
	platformCodeUniqueCheck := make(map[string]struct{})

	for _, platform := range platforms {
		// single-platform check
		if err := checkPlatform(platform); err != nil {
			return err
		}

		// cross-platform check
		if platform.Prefix != "" {
			_, exist := platformPrefixUniqueCheck[platform.Prefix]
			if exist {
				return fmt.Errorf("platform prefix '%v' already exists", platform.Prefix)
			}
			platformPrefixUniqueCheck[platform.Prefix] = struct{}{}
		}
		_, exist := platformCodeUniqueCheck[platform.Code]
		if exist {
			return fmt.Errorf("platform code '%v' already exists", platform.Code)
		}
		platformCodeUniqueCheck[platform.Code] = struct{}{}

		// check platform modules
		modulePrefixUniqueCheck := make(map[string]struct{})
		moduleCodeUniqueCheck := make(map[string]struct{})
		for _, module := range platform.Modules {
			// single-module check
			if err := checkModule(module); err != nil {
				return err
			}

			// cross-module check
			_, exist := modulePrefixUniqueCheck[module.Prefix]
			if exist {
				return fmt.Errorf("module prefix '%v' already exists in platform '%v'", module.Prefix, platform.Name)
			}
			modulePrefixUniqueCheck[module.Prefix] = struct{}{}
			_, exist = moduleCodeUniqueCheck[module.Code]
			if exist {
				return fmt.Errorf("module code '%v' already exists in platform '%v'", module.Code, platform.Name)
			}
			moduleCodeUniqueCheck[module.Code] = struct{}{}

			// check specific error
			for _, specificError := range module.SpecificErrors {
				if err := checkSpecificError(specificError); err != nil {
					return fmt.Errorf("platform: '%v', module: '%v', specificError: '%v' is not valid, %v",
						platform.Name, module.Name, specificError.Suffix, err)
				}

				constantName := platform.Prefix + module.Prefix + specificError.Suffix
				constantCode := platform.Code + module.Code + specificError.Code
				_, err := strconv.ParseInt(constantCode, 10, 32)
				if err != nil {
					return fmt.Errorf("the composed error code of %v cannot be parsed into int32. (platform: %v, module: %v, error suffix: %v)",
						constantName, platform.Name, module.Name, specificError.Suffix)
				}

				_, exist := nameUniqueCheck[constantName]
				if exist {
					return fmt.Errorf("constant name '%v' already exists", constantName)
				}
				_, exist = codeUniqueCheck[constantCode]
				if exist {
					return fmt.Errorf("constant code '%v' already exists", constantCode)
				}

				nameUniqueCheck[constantName] = struct{}{}
				codeUniqueCheck[constantCode] = struct{}{}
			}
		}
	}
	return nil
}

func checkPlatform(platform Platform) error {
	if platform.Code != "10" && platform.Prefix == "" {
		return fmt.Errorf("no prefix is not allow for business related platform: %v", platform.Name)
	}

	if err := checkCode(platform.Code, 2); err != nil {
		return fmt.Errorf("platform '%v' code '%v' is not valid, %v", platform.Name, platform.Code, err)
	}

	num, err := strconv.ParseInt(platform.Code, 10, 32)
	if err != nil {
		return fmt.Errorf("platform '%v' code '%v' is not valid, cannot parse code to int32",
			platform.Name, platform.Code)
	}

	if num != 10 && num < 20 {
		return fmt.Errorf("platform '%v' code '%v' is not valid, business related platform code can not smaller than 20",
			platform.Name, platform.Code)
	}
	return nil
}

func checkModule(module Module) error {
	if module.Prefix == "" {
		return fmt.Errorf("no prefix is not allow for module: %v", module.Name)
	}

	if err := checkCode(module.Code, 3); err != nil {
		return fmt.Errorf("module '%v' code '%v' is not valid, %v", module.Name, module.Code, err)
	}
	return nil
}

func checkSpecificError(specificError SpecificError) error {
	if specificError.Suffix == "" {
		return fmt.Errorf("no suffix is not allow for specific error")
	}

	if err := checkCode(specificError.Code, 3); err != nil {
		return fmt.Errorf("specific error '%v' code '%v' is not valid, %v",
			specificError.Suffix, specificError.Code, err)
	}
	return nil
}

func checkCode(code string, length int) error {
	if len(code) != length {
		return fmt.Errorf("the code length is not equal to %v", length)
	}
	// need to be all digit
	for _, c := range code {
		if c < '0' || c > '9' {
			return fmt.Errorf("the code is not all digit")
		}
	}
	return nil
}
