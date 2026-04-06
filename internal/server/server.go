package server

import (
	"net/http"

	"github.com/RafehMalik/learning-go-shop/internal/config"
	"github.com/RafehMalik/learning-go-shop/internal/services"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"gorm.io/gorm"
)

type Server struct {
	config         *config.Config
	db             *gorm.DB
	logger         zerolog.Logger
	authService    *services.AuthService
	productService *services.ProductService
	userService    *services.UserService
}

func New(cfg *config.Config, db *gorm.DB, logger zerolog.Logger, authService *services.AuthService, productService *services.ProductService, userService *services.UserService) *Server {
	return &Server{
		config:         cfg,
		db:             db,
		logger:         logger,
		authService:    authService,
		productService: productService,
		userService:    userService,
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
	protected := api.Group("/")
	protected.Use(s.authMiddleware())
	{
		user := protected.Group("/user")
		{
			userRoutes := user
			userRoutes.GET("/profile", s.GetProfile)
			userRoutes.PUT("/profile", s.UpdateProfile)

			// category routes
			categories := protected.Group("/categories")
			{
				categoryRoute := categories
				categoryRoute.POST("/", s.adminMiddleware(), s.CreateCategory)
				categoryRoute.PUT("/:id", s.adminMiddleware(), s.UpdateCategory)
				categoryRoute.DELETE("/:id", s.adminMiddleware(), s.DeleteCategory)
			}

			// product routes
			products := protected.Group("/products")
			{
				productRoutes := products
				productRoutes.POST("/", s.adminMiddleware(), s.CreateProduct)
				productRoutes.PUT("/:id", s.adminMiddleware(), s.UpdateProduct)
				productRoutes.DELETE("/:id", s.adminMiddleware(), s.DeleteProduct)

			}
		}
		// public routes
		api.GET("/categories", s.GetCategory)
		api.GET("/products", s.GetProducts)
		api.GET("/products/:id", s.GetProduct)
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
