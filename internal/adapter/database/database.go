package database

import (
	"context"
	"fmt"
	"log/slog"
	"time"

	"github.com/jfelipearaujo-org/ms-product-catalog/internal/environment"
	"github.com/jfelipearaujo-org/ms-product-catalog/internal/shared/health"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type DatabaseService interface {
	GetInstance() *mongo.Database
	health.HealthCheck
}

type Service struct {
	Client *mongo.Client
	db     *mongo.Database
}

func NewDatabase(config *environment.Config) DatabaseService {
	dbAddr := fmt.Sprintf("mongodb://%s:%s@%s:%d/?maxPoolSize=50",
		config.DbConfig.User,
		config.DbConfig.Password,
		config.DbConfig.Host,
		config.DbConfig.Port)

	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(dbAddr))
	if err != nil {
		slog.Error("error connecting to database", "error", err)
		panic(err)
	}

	db := client.Database(config.DbConfig.DbName)

	return &Service{
		Client: client,
		db:     db,
	}
}

func (s *Service) GetInstance() *mongo.Database {
	return s.db
}

func (s *Service) Health() *health.HealthStatus {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	if err := s.Client.Ping(ctx, nil); err != nil {
		slog.Error("could not ping the database", "error", err)
		return &health.HealthStatus{
			Status: "unhealthy",
			Err:    err.Error(),
		}
	}

	return &health.HealthStatus{
		Status: "healthy",
	}
}
