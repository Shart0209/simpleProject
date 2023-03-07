package externalserver

import (
	"context"
	"github.com/jackc/pgx/v5/pgconn"
	"net/http"
	"simpleProject/pkg/model"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
)

type Service interface {
	GetAll() ([]*model.DocumentManagement, error)
	GetByID(uint64) (*model.DocumentManagement, error)
	Create(*model.BindForm) error
	Update(int, *model.BindForm) error
	Delete(uint64) (pgconn.CommandTag, error)
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
	add := &addTransport{
		svc: s.svc,
		log: s.logger.With().Str("transport", "create document").Logger(),
	}
	upd := &updTransport{
		svc: s.svc,
		log: s.logger.With().Str("transport", "update document").Logger(),
	}

	get := &getTransport{
		svc: s.svc,
		log: s.logger.With().Str("transport", "get all/id document").Logger(),
	}

	del := &delTransport{
		svc: s.svc,
		log: s.logger.With().Str("transport", "delete by id document").Logger(),
	}

	doc := s.router.Group("/documents")
	doc.GET("/", s.middleware(get))
	doc.GET("/:id", s.middleware(get))
	doc.POST("/add", s.middleware(add))
	doc.PATCH("/update/:id", s.middleware(upd))
	doc.DELETE("/delete/:id", s.middleware(del))
}

func (s *server) middleware(tr transport) func(*gin.Context) {
	return func(ctx *gin.Context) {
		tr.handler(ctx)
	}
}
