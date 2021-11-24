package generator

import (
	"fmt"
	"os"
)

// EnsureDirExist creates a directory if it does not exist.
func EnsureDirExist(dir string) error {
	err := os.MkdirAll(dir, os.ModePerm)
	if err != nil {
		return err
	}
	return nil
}

// CheckCode check whether the code is valid:
// 1. the code should be in specific length
// 2. the code need to be all digit
func CheckCode(code string, length int) error {
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
