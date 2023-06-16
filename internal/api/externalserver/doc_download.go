package externalserver

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"io"
	"net/http"
	"path/filepath"
)

type getDownloadFileTransport struct {
	svc Service
	log zerolog.Logger
}

func (t *getDownloadFileTransport) Handler(ctx *gin.Context) {
	if fileName := ctx.Param("id"); fileName != "" {
		_, err := filepath.Abs("upload")
		if err != nil {
			t.log.Error().Err(err)
			ctx.AbortWithStatusJSON(http.StatusBadRequest, Response{Error: "path folder to ./upload not found"})
			return
		}

		body, err := io.ReadAll(ctx.Request.Body)
		if err != nil {
			t.log.Error().Err(err)
			ctx.AbortWithStatusJSON(http.StatusBadRequest, Response{Error: "ID from body not found"})
			return
		}

		token := ctx.GetHeader("Authorization")
		if _, err := t.svc.VerifyJWT(token, false); err != nil {
			t.log.Error().Err(err)
			ctx.AbortWithStatusJSON(http.StatusBadRequest, Response{Error: err.Error()})
			return
		}

		var idx string
		err = json.Unmarshal(body, &idx)
		if err != nil {
			t.log.Error().Err(err)
			ctx.AbortWithStatusJSON(http.StatusBadRequest, Response{Error: err.Error()})
		}

		filePath, err := t.svc.GetFilePath(fileName, idx)
		if err != nil {
			t.log.Error().Err(err)
			ctx.AbortWithStatusJSON(http.StatusBadRequest, Response{Error: "file does not exist"})
		}

		ctx.FileAttachment(filepath.Join(filePath, fileName), fileName)
	}
}
