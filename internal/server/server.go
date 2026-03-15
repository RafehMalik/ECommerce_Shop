package server

import (
	"net/http"

	"github.com/RafehMalik/learning-go-shop/internal/config"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"gorm.io/gorm"
)

type Server struct {
	config *config.Config
	db     *gorm.DB
	logger zerolog.Logger
}

func New(cfg *config.Config, db *gorm.DB, logger zerolog.Logger) *Server {
	return &Server{
		config: cfg,
		db:     db,
		logger: logger,
	}
}

func (s *Server) SetupRoutes() *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.Use(s.corsMiddlewares())

	router.GET("/health", s.healthcheck)

	api := router.Group("/api/v1")
	auth := api.Group("/auth")
	{
		auth.POST("/register", s.register)
		auth.POST("/login", s.login)
		auth.POST("/logout", s.logout)
		auth.POST("/refresh", s.refreshToken)
	}
	return router
}

func (s *Server) healthcheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"success": "true"})
}

func (s *Server) corsMiddlewares() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
