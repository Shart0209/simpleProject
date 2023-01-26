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
		idx, err := strconv.Atoi(id)
		if err != nil {
			t.log.Error().Err(err).Msg("bad index")
			ctx.AbortWithStatusJSON(http.StatusBadRequest, Response{Error: "bad request error"})
			return
		}

		data, err := t.svc.GetID(idx)
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
		t.log.Fatal().Err(err).Msg("error get all documents")
		ctx.AbortWithStatusJSON(http.StatusBadRequest, Response{Error: "error not found"})
		return
	}
	ctx.JSON(http.StatusOK, Response{Data: data})
}
