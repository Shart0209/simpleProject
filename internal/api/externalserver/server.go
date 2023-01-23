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
	GetID(id int) (model.DocumentManagement, error)

	Add([]byte) (model.DocumentManagement, error)

	DeleteID(id int) (string, error)
	UpdateID(id int) (string, error)
}

type Server interface {
	Init(bindAddr string)
	SetService(svc Service)
	GetServer() *http.Server
}

type transport interface {
	getAllHandler(ctx *gin.Context)
	getIDHandler(ctx *gin.Context)
	addHandler(ctx *gin.Context)
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
	docm := &docManagementTransport{
		svc: s.svc,
		log: s.logger.With().Str("transport", "Doc").Logger(),
	}

	doc := s.router.Group("/documents")
	doc.GET("/list", s.oneMiddleware(docm))
	doc.GET("/id/:id", s.twoMiddleware(docm))
	doc.POST("/add", s.threeMiddleware(docm))
	//doc.DELETE("/delete/:id", s.middleware())
	//doc.PATCH("/update/:id", s.middleware())

}

func (s *server) oneMiddleware(tr transport) func(*gin.Context) {
	return func(ctx *gin.Context) {
		tr.getAllHandler(ctx)
	}
}

func (s *server) twoMiddleware(tr transport) func(*gin.Context) {
	return func(ctx *gin.Context) {
		tr.getIDHandler(ctx)
	}
}

func (s *server) threeMiddleware(tr transport) func(*gin.Context) {
	return func(ctx *gin.Context) {
		tr.addHandler(ctx)
	}
}
