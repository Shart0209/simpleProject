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

	// get by id document /documents/id/:id
	if id := ctx.Param("id"); id != "" {
		idx, err := strconv.Atoi(id)
		if err != nil {
			t.log.Error().Err(err).Msg("bad index")
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": "bad request error",
			})
			return
		}

		data, err := t.svc.GetID(idx)
		if err != nil {
			t.log.Fatal().Err(err).Msg("error get all json")
		}
		ctx.JSON(http.StatusOK, data)
		return
	}

	// get all documents /documents/list
	data, err := t.svc.GetAll()
	if err != nil {
		t.log.Fatal().Err(err).Msg("error get all json")
	}
	ctx.JSON(http.StatusOK, data)
}
