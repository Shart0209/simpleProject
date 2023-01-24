package externalserver

import (
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"io"
	"net/http"
	"strconv"
)

type docManagementTransport struct {
	svc Service
	log zerolog.Logger
}

func (t *docManagementTransport) addHandler(ctx *gin.Context) {

	resBody, _ := io.ReadAll(ctx.Request.Body)

	data, err := t.svc.Add(resBody)
	if err != nil {
		t.log.Error().Err(err)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, err)
	}
	ctx.JSON(http.StatusOK, data)
}

func (t *docManagementTransport) getAllHandler(ctx *gin.Context) {

	data, err := t.svc.GetAll()
	if err != nil {
		t.log.Fatal().Err(err).Msg("error get all json")
	}
	ctx.JSON(http.StatusOK, data)
}

func (t *docManagementTransport) getIDHandler(ctx *gin.Context) {
	id := ctx.Param("id")
	idd, _ := strconv.Atoi(id)
	data, err := t.svc.GetID(idd)
	if err != nil {
		t.log.Fatal().Err(err).Msg("error get ID not found")
	}
	ctx.JSON(http.StatusOK, data)
}
