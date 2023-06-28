package externalserver

import (
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"net/http"
)

type refreshTransport struct {
	svc Service
	log zerolog.Logger
}

func (t *refreshTransport) Handler(ctx *gin.Context) {
	oldToken := ctx.GetHeader("Authorization")
	token, err := t.svc.VerifyJWT(oldToken, true)
	if err != nil {
		t.log.Error().Err(err)
	}

	data, err := t.svc.ParseJWT(token)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, Response{Error: "Parse JWT failed"})
		return
	}

	// Check for username and password match, usually from a database
	items, err := t.svc.CheckAuthRefresh(data)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, Response{Error: err.Error()})
		return
	}

	newToken, err := t.svc.GenerateJWT(items)
	if err != nil {
		t.log.Error().Err(err).Send()
		ctx.AbortWithStatusJSON(http.StatusBadRequest, Response{Error: "Generate JWT failed"})
		return
	}

	ctx.JSON(http.StatusOK, LoginResponse{AccessToken: newToken})
}
