package externalserver

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"net/http"
	"strconv"
)

type delTransport struct {
	svc Service
	log zerolog.Logger
}

func (t *delTransport) Handler(ctx *gin.Context) {
	//delete by id document http://apiV1/docs/delete/:id
	id := ctx.Param("id")
	if id == "" {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, Response{Error: "bad request by ID {_} not found"})
	}

	idx, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		t.log.Error().Err(err).Msg("bad index")
		ctx.AbortWithStatusJSON(http.StatusBadRequest, Response{Error: fmt.Sprintf("bad request by ID {%s} not integer", id)})
		return
	}

	err = t.svc.Delete(idx)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, Response{Error: fmt.Sprintf("bad request by ID {%s} not found", id)})
		return
	}

	ctx.AbortWithStatusJSON(http.StatusOK, Response{Data: "Success"})
}
