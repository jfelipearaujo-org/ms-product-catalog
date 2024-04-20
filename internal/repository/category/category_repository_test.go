package category

import (
	"context"
	"testing"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/integration/mtest"

	"github.com/google/uuid"
	"github.com/jfelipearaujo-org/ms-product-catalog/internal/entity"
	"github.com/stretchr/testify/assert"
)

var (
	// A key `ok` and value -1 is considered a mongo error
	mongoInvalidOperation primitive.D = bson.D{
		{Key: "ok", Value: -1},
	}

	mongoCommandError error = mongo.CommandError{
		Message: "command failed",
	}
)

func TestNewCategoryRepository(t *testing.T) {
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))

	mt.Run("Should return a new category repository", func(mt *mtest.T) {
		// Arrange
		expected := &CategoryRepository{
			collection: mt.DB.Collection(CategoryCollection),
		}

		// Act
		resp := NewCategoryRepository(mt.DB)

		// Assert
		assert.IsType(mt, expected, resp)
	})
}

func TestCategoryRepository_Create(t *testing.T) {
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))

	mt.Run("Should create a new category", func(mt *mtest.T) {
		// Arrange
		repo := NewCategoryRepository(mt.DB)

		mt.AddMockResponses(mtest.CreateSuccessResponse())

		now := time.Now()

		category := entity.NewCategory("Category 1", "Category 1 description", now)

		// Act
		err := repo.Create(context.Background(), category)

		// Assert
		assert.NoError(mt, err)
	})
}

func TestCategoryRepository_GetByID(t *testing.T) {
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))

	mt.Run("Should get a category by ID", func(mt *mtest.T) {
		// Arrange
		repo := NewCategoryRepository(mt.DB)

		id := uuid.NewString()
		now := time.Now().UTC().Truncate(time.Millisecond)

		mt.AddMockResponses(mtest.CreateCursorResponse(1,
			"foo.bar",
			mtest.FirstBatch,
			bson.D{
				{Key: "uuid", Value: id},
				{Key: "title", Value: "Category 1"},
				{Key: "description", Value: "Category 1 description"},
				{Key: "created_at", Value: now},
				{Key: "updated_at", Value: now},
			}))

		expected := &entity.Category{
			UUID:        id,
			Title:       "Category 1",
			Description: "Category 1 description",
			CreatedAt:   now,
			UpdatedAt:   now,
		}

		// Act
		resp, err := repo.GetByID(context.Background(), id)

		// Assert
		assert.NoError(mt, err)
		assert.Equal(mt, expected, resp)
	})

	mt.Run("Should return error when Decode fails", func(mt *mtest.T) {
		// Arrange
		repo := NewCategoryRepository(mt.DB)

		id := uuid.NewString()

		mt.AddMockResponses(mongoInvalidOperation)

		// Act
		resp, err := repo.GetByID(context.Background(), id)

		// Assert
		assert.Error(mt, err)
		assert.Nil(mt, resp)
	})

	mt.Run("Should return error when nothing is found", func(mt *mtest.T) {
		// Arrange
		repo := NewCategoryRepository(mt.DB)

		id := uuid.NewString()

		mt.AddMockResponses(mtest.CreateCursorResponse(0,
			"foo.bar",
			mtest.FirstBatch))

		// Act
		resp, err := repo.GetByID(context.Background(), id)

		// Assert
		assert.Error(mt, err)
		assert.ErrorIs(mt, err, ErrCategoryNotFound)
		assert.Nil(mt, resp)
	})
}

func TestCategoryRepository_GetByTitle(t *testing.T) {
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))

	mt.Run("Should get a category by Title", func(mt *mtest.T) {
		// Arrange
		repo := NewCategoryRepository(mt.DB)

		id := uuid.NewString()
		now := time.Now().UTC().Truncate(time.Millisecond)

		mt.AddMockResponses(mtest.CreateCursorResponse(1,
			"foo.bar",
			mtest.FirstBatch,
			bson.D{
				{Key: "uuid", Value: id},
				{Key: "title", Value: "Category 1"},
				{Key: "description", Value: "Category 1 description"},
				{Key: "created_at", Value: now},
				{Key: "updated_at", Value: now},
			}))

		expected := &entity.Category{
			UUID:        id,
			Title:       "Category 1",
			Description: "Category 1 description",
			CreatedAt:   now,
			UpdatedAt:   now,
		}

		// Act
		resp, err := repo.GetByTitle(context.Background(), "Category 1")

		// Assert
		assert.NoError(mt, err)
		assert.Equal(mt, expected, resp)
	})

	mt.Run("Should return error when Decode fails", func(mt *mtest.T) {
		// Arrange
		repo := NewCategoryRepository(mt.DB)

		mt.AddMockResponses(mongoInvalidOperation)

		// Act
		resp, err := repo.GetByTitle(context.Background(), "Category 1")

		// Assert
		assert.Error(mt, err)
		assert.Nil(mt, resp)
	})

	mt.Run("Should return error when nothing is found", func(mt *mtest.T) {
		// Arrange
		repo := NewCategoryRepository(mt.DB)

		mt.AddMockResponses(mtest.CreateCursorResponse(0,
			"foo.bar",
			mtest.FirstBatch))

		// Act
		resp, err := repo.GetByTitle(context.Background(), "Category 1")

		// Assert
		assert.Error(mt, err)
		assert.ErrorIs(mt, err, ErrCategoryNotFound)
		assert.Nil(mt, resp)
	})
}
