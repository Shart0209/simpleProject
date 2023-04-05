package externalserver

import (
	"context"
	"github.com/gin-contrib/cors"
	"net/http"
	"simpleProject/pkg/model"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
)

type Service interface {
	GetAll() ([]*model.DocsAttrs, error)
	GetByID(uint64) (*model.DocsAttrs, error)
	Create(*model.BindForm) error
	Update(int, *model.BindForm) error
	Delete(uint64) error
	GetSps() (*model.Sps, error)
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
	//s.router.Use(cors.New(cors.Config{
	//	AllowOrigins: []string{"*"},
	//}))

	s.router.Use(cors.Default())

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

	sps := &getCatTransport{
		svc: s.svc,
		log: s.logger.With().Str("transport", "get list categories/distributors").Logger(),
	}

	del := &delTransport{
		svc: s.svc,
		log: s.logger.With().Str("transport", "delete by id document").Logger(),
	}

	doc := s.router.Group("/docs")
	doc.GET("/", s.middleware(get))
	doc.GET("/:id", s.middleware(get))
	doc.GET("/sps", s.middleware(sps))
	doc.POST("/add", s.middleware(add))
	doc.PATCH("/update/:id", s.middleware(upd))
	doc.DELETE("/delete/:id", s.middleware(del))
}

func (s *server) middleware(tr transport) func(*gin.Context) {
	return func(ctx *gin.Context) {
		tr.handler(ctx)
	}
}
