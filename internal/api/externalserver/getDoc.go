package externalserver

import (
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"net/http"
	"strconv"
)

type getTransport struct {
	svc Service
	log zerolog.Logger
}

func (t *getTransport) handler(ctx *gin.Context) {

	// get by id document /documents/:id
	if id := ctx.Param("id"); id != "" {
		idx, err := strconv.ParseUint(id, 10, 64)
		if err != nil {
			t.log.Error().Err(err).Msg("bad index")
			ctx.AbortWithStatusJSON(http.StatusBadRequest, Response{Error: "bad request errors"})
			return
		}

		data, err := t.svc.GetByID(idx)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, Response{Error: err.Error()})
			return
		}
		ctx.JSON(http.StatusOK, Response{Data: data})
		return
	}

	// get all documents /documents
	data, err := t.svc.GetAll()
	if err != nil {
		t.log.Fatal().Err(err).Msg("errors get all documents")
		ctx.AbortWithStatusJSON(http.StatusBadRequest, Response{Error: "errors not found"})
		return
	}
	ctx.JSON(http.StatusOK, Response{Data: data})
}
