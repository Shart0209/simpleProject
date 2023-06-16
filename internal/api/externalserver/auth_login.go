package externalserver

import (
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"net/http"
	"strings"
)

type loginTransport struct {
	svc Service
	log zerolog.Logger
}

func (t *loginTransport) Handler(ctx *gin.Context) {
	var loginForm LoginRequest

	if err := ctx.ShouldBind(&loginForm); err != nil {
		t.log.Error().Err(err).Send()
		ctx.AbortWithStatusJSON(http.StatusBadRequest, Response{Error: err.Error()})
		return
	}

	// Validate form input
	if strings.Trim(loginForm.Login, " ") == "" || strings.Trim(loginForm.Password, " ") == "" {
		t.log.Error().Msg("fields username OR password is empty")
		ctx.AbortWithStatusJSON(http.StatusBadRequest, Response{Error: "Parameters can't be empty"})
		return
	}

	// TEST: Check for username and password match, usually from a database
	if err := testCheckAuth(loginForm.Login, loginForm.Password); err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, Response{Error: "Authentication failed"})
		return
	}

	token, err := t.svc.GenerateJWT(&loginForm.Login)
	if err != nil {
		t.log.Error().Err(err).Send()
		ctx.AbortWithStatusJSON(http.StatusBadRequest, Response{Error: "Generate JWT failed"})
		return
	}

	ctx.JSON(http.StatusOK, LoginResponse{AccessToken: token})

}
