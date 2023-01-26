package externalserver

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/rs/zerolog"
	"path/filepath"
)

func upload(ctx *gin.Context, t *zerolog.Logger) ([]string, error) {

	form, err := ctx.MultipartForm()
	if err != nil {
		return nil, fmt.Errorf("get form err: %s", err.Error())
	}

	files := form.File["files"]

	dirName := "./upload"
	baseDir, err := filepath.Abs(dirName)
	if err != nil {
		return nil, fmt.Errorf("path dir to ./upload not found")
	}
	arr := make([]string, 0)
	for _, file := range files {
		fname := filepath.Base(file.Filename)
		arr = append(arr, fname)
		filename := baseDir + `/` + uuid.New().String() + filepath.Ext(fname)
		if err := ctx.SaveUploadedFile(file, filename); err != nil {
			return nil, fmt.Errorf("upload file err: %s", err.Error())
		}
	}
	return arr, fmt.Errorf("uploaded successfully")
}
