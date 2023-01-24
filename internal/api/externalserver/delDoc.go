package externalserver

import (
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
)

type delTransport struct {
	svc Service
	log zerolog.Logger
}

func (t *delTransport) handler(ctx *gin.Context) {}
