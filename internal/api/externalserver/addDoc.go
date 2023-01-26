package externalserver

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"io"
	"net/http"
)

type addTransport struct {
	svc Service
	log zerolog.Logger
}

func (t *addTransport) handler(ctx *gin.Context) {

	var body interface{}
	tmp := ctx.BindJSON(&body)
	fmt.Println(tmp)

	resBody, err := io.ReadAll(ctx.Request.Body)
	if err != nil || len(resBody) == 0 {
		t.log.Error().Err(err).Msg("bad read context body")
		ctx.AbortWithStatusJSON(http.StatusBadRequest, Response{Error: err})
		return
	}

	arr, err := upload(ctx, &t.log)
	if err != nil {
		t.log.Error().Err(err)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, Response{Error: err})
		return
	}

	// add document /documents/add
	data, err := t.svc.Add(&resBody, &arr)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, Response{Error: err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, Response{Data: data})
}
