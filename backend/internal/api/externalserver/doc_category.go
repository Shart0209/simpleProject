package externalserver

import (
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"net/http"
)

type getCategoryTransport struct {
	svc Service
	log zerolog.Logger
}

func (t *getCategoryTransport) Handler(ctx *gin.Context) {

	// get list categories/distributor http://apiV1/docs/sps
	data, err := t.svc.GetSps()
	if err != nil {
		t.log.Fatal().Err(err).Msg("errors get list categories/distributor")
		ctx.AbortWithStatusJSON(http.StatusBadRequest, Response{Error: "errors not found"})
		return
	}
	ctx.JSON(http.StatusOK, Response{Data: data})
}
