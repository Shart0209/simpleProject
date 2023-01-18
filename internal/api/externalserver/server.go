package externalserver

import (
	"context"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"net/http"
)

type Server interface {
	Init(addr string)
	SetService(svc string)
	GetServer() *http.Server
}

type server struct {
	ctx    context.Context
	svc    string
	router *gin.Engine
	http   *http.Server
	logger zerolog.Logger
}

func (s *server) SetService(svc string) {
	s.svc = svc
	s.configureRouter()
}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

func New(ctx context.Context, logger zerolog.Logger) *server {
	s := &server{
		router: gin.New(),
		logger: logger,
		ctx:    ctx,
	}
	return s
}

func (s *server) Init(addr string) {
	s.http = &http.Server{
		Addr:    addr,
		Handler: s,
	}
}

func (s *server) GetServer() *http.Server {
	return s.http
}

func (s *server) configureRouter() {
	s.router.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
	}))

	s.router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, "pong")
	})

}
