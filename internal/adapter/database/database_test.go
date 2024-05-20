package database

import (
	"testing"

	"github.com/jfelipearaujo-org/ms-product-catalog/internal/environment"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/mongo/integration/mtest"
)

func TestGetInstance(t *testing.T) {
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))

	mt.Run("Should return healthy status", func(mt *mtest.T) {
		// Arrange
		config := &environment.Config{
			DbConfig: &environment.DatabaseConfig{
				DbName:        "db",
				Url:           "mongodb://host:1234",
				UrlSecretName: "db-secret-url",
			},
		}

		mt.AddMockResponses(mtest.CreateSuccessResponse())

		service := NewDatabase(config)
		service.(*Service).Client = mt.Client

		// Act
		res := service.GetInstance()

		// Assert
		assert.NotNil(mt, res)
	})
}

func TestHealth(t *testing.T) {
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))

	mt.Run("Should return healthy status", func(mt *mtest.T) {
		// Arrange
		config := &environment.Config{
			DbConfig: &environment.DatabaseConfig{
				DbName:        "db",
				Url:           "mongodb://host:1234",
				UrlSecretName: "db-secret-url",
			},
		}

		mt.AddMockResponses(mtest.CreateSuccessResponse())

		service := NewDatabase(config)
		service.(*Service).Client = mt.Client

		// Act
		res := service.Health()

		// Assert
		assert.NotNil(mt, res)
		assert.Equal(mt, "healthy", res.Status)
	})

	mt.Run("Should return healthy status", func(mt *mtest.T) {
		// Arrange
		config := &environment.Config{
			DbConfig: &environment.DatabaseConfig{
				DbName:        "db",
				Url:           "mongodb://host:1234",
				UrlSecretName: "db-secret-url",
			},
		}

		service := NewDatabase(config)
		service.(*Service).Client = mt.Client

		// Act
		res := service.Health()

		// Assert
		assert.NotNil(mt, res)
		assert.Equal(mt, "unhealthy", res.Status)
	})
}
