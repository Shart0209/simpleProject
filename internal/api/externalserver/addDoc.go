package externalserver

import (
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
)

type addTransport struct {
	svc Service
	log zerolog.Logger
}

func (t *addTransport) handler(ctx *gin.Context) {}
