package externalserver

import (
	"context"
	"net/http"
	"simpleProject/pkg/model"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
)

type Service interface {
	GetAll() (map[int]model.DocumentManagement, error)
	Add() (string, error)
	DeleteID(id int) (string, error)
	GetID(id int) (string, error)
	UpdateID(id int) (string, error)
}

type Server interface {
	Init(bindAddr string)
	SetService(svc Service)
	GetServer() *http.Server
}

type transport interface {
	handler(ctx *gin.Context)
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

	//test
	s.router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	// init transports
	docManagement := &docManagementTransport{
		svc: s.svc,
		log: s.logger.With().Str("transport", "docManagement").Logger(),
	}

	docAdd := &docAddTransport{
		svc: s.svc,
		log: s.logger.With().Str("transport", "docManagement").Logger(),
	}

	//#######################################
	//	добавить handlers на каждый метод
	//  создать файлики под каждый handler
	//  добавить свой транспорт для каждой ручки
	//#######################################

	doc := s.router.Group("/documents")
	doc.GET("/list", s.middleware())
	doc.GET("/id/:id", s.middleware())
	doc.POST("/add", s.middleware())
	doc.DELETE("/delete/:id", s.middleware())
	doc.PATCH("/update/:id", s.middleware())

}

func (s *server) middleware(tr transport) func(*gin.Context) {
	return func(ctx *gin.Context) {
		tr.handler(ctx)
	}
}
