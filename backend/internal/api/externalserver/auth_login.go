package externalserver

import (
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"net/http"
	"simpleProject/pkg/model"
	"strings"
)

type authTransport struct {
	svc Service
	log zerolog.Logger
}

func (t *authTransport) Handler(ctx *gin.Context) {
	var loginForm model.LoginRequest

	if err := ctx.ShouldBind(&loginForm); err != nil {
		t.log.Error().Err(err).Send()
		ctx.AbortWithStatusJSON(http.StatusBadRequest, Response{Error: "Login or password incorrect"})
		return
	}

	// Validate form input
	if strings.Trim(loginForm.Login, " ") == "" || strings.Trim(loginForm.Password, " ") == "" {
		t.log.Error().Msg("fields username OR password is empty")
		ctx.AbortWithStatusJSON(http.StatusBadRequest, Response{Error: "Parameters can't be empty"})
		return
	}

	// Check for username and password match, usually from a database
	data, err := t.svc.CheckAuthLogin(&loginForm)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, Response{Error: err.Error()})
		return
	}

	token, err := t.svc.GenerateJWT(data)
	if err != nil {
		t.log.Error().Err(err).Send()
		ctx.AbortWithStatusJSON(http.StatusBadRequest, Response{Error: "Generate JWT failed"})
		return
	}

	ctx.JSON(http.StatusOK, LoginResponse{AccessToken: token})

}
