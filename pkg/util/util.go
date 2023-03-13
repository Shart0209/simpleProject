package util

import (
	"fmt"
	"github.com/google/uuid"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
	"simpleProject/pkg/model"
	"strconv"
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

func DeleteFile(path *string) error {

	//checking that the file does not exist

	if _, err := os.Stat(*path); err != nil {
		if os.IsNotExist(err) {
			fmt.Println("file does not exist")
			return err
		}
	}
	err := os.Remove(*path)
	if err != nil {
		return err
	}

	return nil
}

func ParserBindForm(bindFiles []*multipart.FileHeader, folder string, data *[]model.Files) error {

	fileID := uuid.New().String()

	baseDir, err := filepath.Abs(folder)
	if err != nil {
		return fmt.Errorf("path folder to ./upload not found")
	}

	for i, file := range bindFiles {
		ext := filepath.Base(file.Filename)
		filename := fileID + "-" + strconv.Itoa(i) + filepath.Ext(ext)
		filePath := filepath.Join(baseDir, filename)

		tmp := model.Files{
			File: model.File{
				Name: filename,
				Size: int(file.Size),
				Path: filePath,
			},
			Files: file,
		}

		*data = append(*data, tmp)
	}

	return nil
}

func SaveUploadedFile(file *multipart.FileHeader, dst *string) error {

	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	out, err := os.Create(*dst)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, src)
	return err
}
