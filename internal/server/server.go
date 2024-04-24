package server

import (
	"fmt"
	"net/http"
	"time"

	"github.com/jfelipearaujo-org/ms-product-catalog/internal/adapter/database"
	"github.com/jfelipearaujo-org/ms-product-catalog/internal/environment"
	"github.com/jfelipearaujo-org/ms-product-catalog/internal/handler/category/create_category_handler"
	"github.com/jfelipearaujo-org/ms-product-catalog/internal/handler/category/delete_category_handler"
	"github.com/jfelipearaujo-org/ms-product-catalog/internal/handler/category/get_categories_handler"
	"github.com/jfelipearaujo-org/ms-product-catalog/internal/handler/category/get_category_handler"
	"github.com/jfelipearaujo-org/ms-product-catalog/internal/handler/health_handler"
	"github.com/jfelipearaujo-org/ms-product-catalog/internal/provider/time_provider"
	category_repository "github.com/jfelipearaujo-org/ms-product-catalog/internal/repository/category"
	"github.com/jfelipearaujo-org/ms-product-catalog/internal/service/category/create_category"
	"github.com/jfelipearaujo-org/ms-product-catalog/internal/service/category/delete_category"
	"github.com/jfelipearaujo-org/ms-product-catalog/internal/service/category/get_categories"
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
	timeProvider := time_provider.NewTimeProvider(time.Now)

	// repositories
	categoryRepository := category_repository.NewRepository(server.db.GetInstance())

	// services
	getCategoryService := get_category.NewService(categoryRepository)
	getCategoriesService := get_categories.NewService(categoryRepository)
	createCategoryService := create_category.NewService(categoryRepository, timeProvider)
	deleteCategoryService := delete_category.NewService(categoryRepository)

	// handlers
	getCategoryHandler := get_category_handler.NewHandler(getCategoryService)
	getCategoriesHandler := get_categories_handler.NewHandler(getCategoriesService)
	createCategoryHandler := create_category_handler.NewHandler(createCategoryService, getCategoriesService)
	deleteCategoryHandler := delete_category_handler.NewHandler(deleteCategoryService)

	// routes
	group.GET("/categories", getCategoriesHandler.Handle)
	group.GET("/categories/:id", getCategoryHandler.Handle)
	group.POST("/categories", createCategoryHandler.Handle)
	group.DELETE("/categories/:id", deleteCategoryHandler.Handle)
}
