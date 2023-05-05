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
		log: s.logger.With().Str("transport", "create").Logger(),
	}
	upd := &updTransport{
		svc: s.svc,
		log: s.logger.With().Str("transport", "update").Logger(),
	}

	get := &getTransport{
		svc: s.svc,
		log: s.logger.With().Str("transport", "get all/id").Logger(),
	}

	download := &getDownloadFileTransport{
		svc: s.svc,
		log: s.logger.With().Str("transport", "get file").Logger(),
	}

	sps := &getCatTransport{
		svc: s.svc,
		log: s.logger.With().Str("transport", "get list categories/distributors").Logger(),
	}

	del := &delTransport{
		svc: s.svc,
		log: s.logger.With().Str("transport", "delete by id").Logger(),
	}

	api := s.router.Group("api/docs")
	api.GET("/", s.middleware(get))
	api.GET("/:id", s.middleware(get))
	api.GET("/sps", s.middleware(sps))
	api.GET("/download/:id", s.middleware(download))
	api.POST("/add", s.middleware(add))
	api.PATCH("/update/:id", s.middleware(upd))
	api.DELETE("/delete/:id", s.middleware(del))
}

func (s *server) middleware(tr transport) func(*gin.Context) {
	return func(ctx *gin.Context) {
		tr.handler(ctx)
	}
}
