package util

import (
	"fmt"
	"os"
	"path/filepath"
)

func CreateFolder(path *string) error {

	_, err := os.Stat(*path)
	if err != nil && os.IsNotExist(err) {
		err = os.MkdirAll(*path, os.ModePerm)
		if err != nil {
			return err
		}
	}
	return nil
}

func DeleteFile(arr *[]string, path *string) error {

	baseDir, err := filepath.Abs(*path)
	if err != nil {
		return fmt.Errorf("path folder to ./upload not found")
	}

	if len(*arr) == 0 {
		readDir, _ := os.Open(baseDir)
		files, _ := readDir.Readdir(0)

		for f := range files {
			file := files[f]

			fileName := file.Name()
			filePath := filepath.Join(baseDir, fileName)

			_ = os.Remove(filePath)
		}
	}

	return nil
}
