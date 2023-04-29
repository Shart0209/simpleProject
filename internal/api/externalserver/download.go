package externalserver

import (
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"net/http"
	"path/filepath"
)

type getDownloadFileTransport struct {
	svc Service
	log zerolog.Logger
}

func (t *getDownloadFileTransport) handler(ctx *gin.Context) {
	if fileName := ctx.Param("id"); fileName != "" {
		path, err := filepath.Abs("upload")
		if err != nil {
			t.log.Error().Err(err)
			ctx.AbortWithStatusJSON(http.StatusBadRequest, Response{Error: "path folder to ./upload not found"})
			return
		}

		ctx.FileAttachment(path+"/"+fileName, fileName)
	}
}
