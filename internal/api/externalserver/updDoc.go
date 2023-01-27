package externalserver

import (
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"net/http"
	"simpleProject/pkg/model"
	"strconv"
)

type updTransport struct {
	svc Service
	log zerolog.Logger
}

func (t *updTransport) handler(ctx *gin.Context) {
	var bindForm model.BindForm
	if err := ctx.ShouldBind(&bindForm); err != nil {
		t.log.Error().Err(err).Send()

		ctx.AbortWithStatusJSON(http.StatusBadRequest, Response{Error: err.Error()})
		return
	}

	id := ctx.Param("id")
	idx, err := strconv.Atoi(id)
	if err != nil {
		t.log.Error().Err(err).Send()
		ctx.AbortWithStatusJSON(http.StatusBadRequest, Response{Error: err.Error()})
		return
	}

	if err := t.svc.UpdateID(idx, &bindForm); err != nil {
		t.log.Error().Err(err).Send()
		ctx.AbortWithStatusJSON(http.StatusBadRequest, Response{Error: err.Error()})
		return
	}

	ctx.Status(http.StatusOK)
}
