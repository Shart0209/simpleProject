package externalserver

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
)

type docManagmentTransport struct {
	svc Service
	log zerolog.Logger
}

func (t *docManagmentTransport) handler(ctx *gin.Context) {
	isError := ctx.Query("is_error")
	if isError != "" {
		t.log.Error().Msg("fake error")

		ctx.AbortWithStatusJSON(http.StatusBadRequest, Response{
			Error: "fake error",
		})

		return
	}

	name := ctx.Query("name")

	fmt.Println(name)

	data, err := t.svc.Doc(name)
	if err != nil {
		t.log.Error().Err(err).Msg("t.svc.HelloWorld")

		ctx.AbortWithStatusJSON(http.StatusInternalServerError, Response{
			Error: err.Error(),
		})

		return
	}

	ctx.JSON(http.StatusOK, data)
}
