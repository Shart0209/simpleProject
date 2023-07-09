package externalserver

import (
	"context"
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/rs/zerolog"
	"net/http"
	"simpleProject/pkg/model"
	"time"
)

type Service interface {
	GetAll() ([]*model.DocsAttrs, error)
	GetByID(uint64) (*model.DocsAttrs, error)
	Create(*model.BindForm) error
	Update(*model.BindForm) error
	Delete(uint64) error
	GetSps() (*model.Sps, error)
	GetFilePath(string, string) (string, error)
	GenerateJWT(*model.Auth) (string, error)
	VerifyJWT(string, bool) (*jwt.Token, error)
	ParseJWT(*jwt.Token) (jwt.MapClaims, error)
	CheckAuthLogin(*model.LoginRequest) (*model.Auth, error)
	CheckAuthRefresh(jwt.MapClaims) (*model.Auth, error)
}

type Server interface {
	Init(bindPort, bindAddr string)
	SetService(svc Service)
	GetServer() *http.Server
}

type transport interface {
	Handler(ctx *gin.Context)
}

type server struct {
	ctx    context.Context
	router *gin.Engine
	logger zerolog.Logger
	svc    Service
	http   *http.Server
}

func New(ctx context.Context, isDebug bool, logger zerolog.Logger) *server {
	if !isDebug {
		gin.SetMode(gin.ReleaseMode)
	}

	s := &server{
		router: gin.New(),
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

func (s *server) Init(bindPort, bindAddr string) {
	s.http = &http.Server{
		Addr:    fmt.Sprintf("%s:%s", bindAddr, bindPort),
		Handler: s,
	}
	s.logger.Info().Msg(s.http.Addr)
}

func (s *server) GetServer() *http.Server {
	return s.http
}

func (s *server) configureRouter() {
	s.router.Use(cors.New(cors.Config{
		AllowOrigins:  []string{"*"},
		AllowMethods:  []string{"POST", "GET", "OPTIONS", "DELETE"},
		AllowHeaders:  []string{"Origin", "Accept", "Authorization", "Content-Type", "Accept-Encoding"},
		ExposeHeaders: []string{"Content-Length", "Access-Control-Allow-Origin", "Access-Control-Allow-Methods", "Access-Control-Allow-Headers"},
		MaxAge:        12 * time.Hour,
	}))

	// logs
	s.router.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		// your custom format
		return fmt.Sprintf("%s - [%s] \"%s %s %s %d %s \"%s\" %s\"\n",
			param.ClientIP,
			param.TimeStamp.Format(time.RFC1123),
			param.Method,
			param.Path,
			param.Request.Proto,
			param.StatusCode,
			param.Latency,
			param.Request.UserAgent(),
			param.ErrorMessage,
		)
	}))

	// init transports
	add := &addTransport{
		svc: s.svc,
		log: s.logger.With().Str("transport", "docs/create").Logger(),
	}
	upd := &updTransport{
		svc: s.svc,
		log: s.logger.With().Str("transport", "docs/update").Logger(),
	}

	get := &getTransport{
		svc: s.svc,
		log: s.logger.With().Str("transport", "docs/all_ID").Logger(),
	}

	download := &getDownloadFileTransport{
		svc: s.svc,
		log: s.logger.With().Str("transport", "docs/downloadFile").Logger(),
	}

	sps := &getCategoryTransport{
		svc: s.svc,
		log: s.logger.With().Str("transport", "docs/categories").Logger(),
	}

	del := &delTransport{
		svc: s.svc,
		log: s.logger.With().Str("transport", "docs/deleteID").Logger(),
	}

	authorize := &authTransport{
		svc: s.svc,
		log: s.logger.With().Str("transport", "auth/login").Logger(),
	}
	refresh := &refreshTransport{
		svc: s.svc,
		log: s.logger.With().Str("transport", "auth/refresh token").Logger(),
	}

	//test
	s.router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	apiV1 := s.router.Group("apiV1/docs")
	apiV1.GET("/", s.middleware(get))
	apiV1.GET("/:id", s.middleware(get))
	apiV1.GET("/sps", s.middleware(sps))
	apiV1.POST("/download/:id", s.authMiddleware, s.middleware(download))
	apiV1.POST("/add", s.authMiddleware, s.middleware(add))
	apiV1.POST("/update/:id", s.authMiddleware, s.middleware(upd))
	apiV1.DELETE("/delete/:id", s.authMiddleware, s.middleware(del))

	apiV2 := s.router.Group("apiV2/auth")
	apiV2.POST("/login", s.middleware(authorize))
	apiV2.POST("/refresh", s.middleware(refresh))

}

func (s *server) authMiddleware(ctx *gin.Context) {
	token := ctx.GetHeader("Authorization")
	if _, err := s.svc.VerifyJWT(token, false); err != nil {
		s.logger.Error().Err(err).Send()
		ctx.AbortWithStatusJSON(http.StatusBadRequest, Response{Error: err.Error()})

		return
	}

	ctx.Next()
}

func (s *server) middleware(tr transport) func(*gin.Context) {
	return func(ctx *gin.Context) {
		tr.Handler(ctx)
	}
}
