package externalserver

import (
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"io"
	"net/http"
	"strconv"
)

type updTransport struct {
	svc Service
	log zerolog.Logger
}

func (t *updTransport) handler(ctx *gin.Context) {
	resBody, err := io.ReadAll(ctx.Request.Body)
	if err != nil {
		t.log.Error().Err(err).Msg("bad read context body")
		ctx.AbortWithStatusJSON(http.StatusBadRequest, Response{Error: err})
		return
	}

	id := ctx.Param("id")
	idx, err := strconv.Atoi(id)
	if err != nil {
		t.log.Error().Err(err).Msg("bad index")
		ctx.AbortWithStatusJSON(http.StatusBadRequest, Response{Error: "bad request error"})
		return
	}

	// update document /documents/update
	data, err := t.svc.UpdateID(idx, &resBody)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, Response{Error: err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, Response{Data: data})

}
