package generator

import "os"

// EnsureDirExist creates a directory if it does not exist.
func EnsureDirExist(dir string) error {
	err := os.MkdirAll(dir, os.ModePerm)
	if err != nil {
		return err
	}
	return nil
}
