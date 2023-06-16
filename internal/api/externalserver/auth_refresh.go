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
		ctx.AbortWithStatusJSON(http.StatusBadRequest, Response{Error: err.Error()})
		return
	}

	login, err := t.svc.ParseJWT(token)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, Response{Error: "Parse JWT failed"})
		return
	}

	// TEST: Check for username and password match, usually from a database
	if login != "test" {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, Response{Error: "JWT login failed"})
		return
	}

	newToken, err := t.svc.GenerateJWT(&login)
	if err != nil {
		t.log.Error().Err(err).Send()
		ctx.AbortWithStatusJSON(http.StatusBadRequest, Response{Error: "Generate JWT failed"})
		return
	}

	ctx.JSON(http.StatusOK, LoginResponse{AccessToken: newToken})
}
