package server

import (
	"fmt"
	"net/http"
	"time"

	"github.com/jfelipearaujo-org/ms-product-catalog/internal/adapter/database"
	"github.com/jfelipearaujo-org/ms-product-catalog/internal/environment"
	"github.com/jfelipearaujo-org/ms-product-catalog/internal/handler/category/get_category_handler"
	"github.com/jfelipearaujo-org/ms-product-catalog/internal/handler/health_handler"
	category_repository "github.com/jfelipearaujo-org/ms-product-catalog/internal/repository/category"
	"github.com/jfelipearaujo-org/ms-product-catalog/internal/service/category/get_category"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Server struct {
	Config *environment.Config
	db     database.DatabaseService
}

func NewServer(config *environment.Config) *http.Server {
	server := &Server{
		Config: config,
		db:     database.NewDatabase(config),
	}

	return &http.Server{
		Addr:         fmt.Sprintf(":%d", server.Config.ApiConfig.Port),
		Handler:      server.RegisterRoutes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}
}

func (server *Server) RegisterRoutes() http.Handler {
	e := echo.New()
	e.Use(middleware.Recover())

	server.registerHealthCheck(e)

	group := e.Group(fmt.Sprintf("/api/%s", server.Config.ApiConfig.ApiVersion))

	server.registerCategoryRoutes(group)

	return e
}

func (server *Server) registerHealthCheck(e *echo.Echo) {
	healthHandler := health_handler.NewHandler(server.db)

	e.GET("/health", healthHandler.Handle)
}

func (server *Server) registerCategoryRoutes(group *echo.Group) {
	// repositories
	categoryRepository := category_repository.NewRepository(server.db.GetInstance())

	// services
	getCategoryService := get_category.NewService(categoryRepository)

	// handlers
	getCategoryHandler := get_category_handler.NewHandler(getCategoryService)

	// routes
	group.GET("/categories", getCategoryHandler.Handle)
}
