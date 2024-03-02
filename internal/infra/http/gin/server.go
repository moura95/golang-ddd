package gin

import (
	"time"

	"github.com/axiaoxin-com/ratelimiter"
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"github.com/moura95/go-ddd/internal/infra/cfg"

	"go.uber.org/zap"
)

type Server struct {
	store  *sqlx.DB
	router *gin.Engine
	config *cfg.Config
	logger *zap.SugaredLogger
}

func NewServer(cfg cfg.Config, store *sqlx.DB, log *zap.SugaredLogger) *Server {

	server := &Server{
		store:  store,
		config: &cfg,
		logger: log,
	}
	var router *gin.Engine

	if server.config.GinMode == "release" {
		router = gin.Default()
		router.Use(gzip.Gzip(gzip.DefaultCompression))
		router.Use(ratelimiter.GinMemRatelimiter(ratelimiter.GinRatelimiterConfig{
			LimitKey: func(c *gin.Context) string {
				return c.ClientIP()
			},
			LimitedHandler: func(c *gin.Context) {
				c.JSON(200, "too many requests!!!")
				c.Abort()
				return
			},
			TokenBucketConfig: func(*gin.Context) (time.Duration, int) {
				return time.Second * 60, 4000
			},
		}))

	} else {
		router = gin.Default()
	}

	corsConfig := cors.Config{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "redirect", "Authorization"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}

	router.Use(cors.New(corsConfig))
	server.createRoutesV1(router, log)

	server.router = router
	return server
}

func (s *Server) Start(address string) error {
	return s.router.Run(address)
}
