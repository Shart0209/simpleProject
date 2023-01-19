package externalserver

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
)

type Service interface {
	Hello(name string) (string, error)
}

type Server interface {
	Init(bindAddr string)
	SetService(svc Service)
	GetServer() *http.Server
}

type transport interface {
	handler(ctx *gin.Context)
	name() string
}

type server struct {
	ctx    context.Context
	router *gin.Engine
	logger zerolog.Logger
	svc    Service
	http   *http.Server
}

func New(ctx context.Context, logger zerolog.Logger) *server {
	s := &server{
		router: gin.Default(),
		logger: logger,
		ctx:    ctx,
	}

	return s
}

func (s *server) SetService(svc Service) {
	s.svc = svc
	s.configureRouter()
}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

func (s *server) Init(bindAddr string) {
	s.http = &http.Server{
		Addr:    bindAddr,
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
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	// init transports
	// hello := &helloTransport{
	// 	svc: s.svc,
	// 	log: s.logger.With().Str("transport", "hello").Logger(),
	// }

	// apiV1 := s.router.Group("/api/v1")
	// apiV1.GET("/hello", s.someMiddleware(hello))

}

func (s *server) someMiddleware(tr transport) func(*gin.Context) {
	return func(ctx *gin.Context) {
		t := time.Now()

		defer func() {
			s.logger.Info().Str("Api call time lead:", time.Since(t).String()).Msg(ctx.Request.RequestURI)
		}()

		tr.handler(ctx)
	}
}
