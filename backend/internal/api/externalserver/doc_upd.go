package externalserver

import (
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"net/http"
	"simpleProject/pkg/model"
)

type updTransport struct {
	svc Service
	log zerolog.Logger
}

func (t *updTransport) Handler(ctx *gin.Context) {

	var bindForm model.BindForm

	if err := ctx.ShouldBind(&bindForm); err != nil {
		t.log.Error().Err(err).Send()
		ctx.AbortWithStatusJSON(http.StatusBadRequest, Response{Error: err.Error()})
		return
	}

	if err := t.svc.Update(&bindForm); err != nil {
		t.log.Error().Err(err).Send()
		ctx.AbortWithStatusJSON(http.StatusBadRequest, Response{Error: err.Error()})
		return
	}

	ctx.Status(http.StatusOK)
}
