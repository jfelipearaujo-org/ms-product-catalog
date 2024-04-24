package category

import (
	"context"
	"testing"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/integration/mtest"

	"github.com/google/uuid"
	"github.com/jfelipearaujo-org/ms-product-catalog/internal/common"
	"github.com/jfelipearaujo-org/ms-product-catalog/internal/entity"
	"github.com/stretchr/testify/assert"
)

var (
	// A key `ok` and value -1 is considered a mongo error
	mongoInvalidOperation primitive.D = bson.D{
		{Key: "ok", Value: -1},
	}
)

func getPrimitiveCategory(id string, now time.Time) bson.D {
	return bson.D{
		{Key: "uuid", Value: id},
		{Key: "title", Value: "category title"},
		{Key: "description", Value: "category description"},
		{Key: "created_at", Value: now},
		{Key: "updated_at", Value: now},
	}
}

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

		emptyId := primitive.NewObjectID()

		category := entity.NewCategory("category title", "category description", now)

		// Act
		err := repo.Create(context.Background(), category)

		// Assert
		assert.NoError(mt, err)
		assert.NotEqual(mt, emptyId, category.Id)
	})

	mt.Run("Should return error when InsertOne fails", func(mt *mtest.T) {
		// Arrange
		repo := NewCategoryRepository(mt.DB)

		mt.AddMockResponses(mongoInvalidOperation)

		now := time.Now()

		category := entity.NewCategory("category title", "category description", now)

		// Act
		err := repo.Create(context.Background(), category)

		// Assert
		assert.Error(mt, err)
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
			getPrimitiveCategory(id, now)))

		expected := entity.Category{
			UUID:        id,
			Title:       "category title",
			Description: "category description",
			CreatedAt:   now,
			UpdatedAt:   now,
		}

		// Act
		resp, err := repo.GetByID(context.Background(), id)

		// Assert
		assert.NoError(mt, err)
		assert.Equal(mt, expected, resp)
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
			getPrimitiveCategory(id, now)))

		expected := entity.Category{
			UUID:        id,
			Title:       "category title",
			Description: "category description",
			CreatedAt:   now,
			UpdatedAt:   now,
		}

		// Act
		resp, err := repo.GetByTitle(context.Background(), "category title")

		// Assert
		assert.NoError(mt, err)
		assert.Equal(mt, expected, resp)
	})
}

func TestCategoryRepository_GetAll(t *testing.T) {
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))

	mt.Run("Should return a list of categories paginated", func(mt *mtest.T) {
		// Arrange
		repo := NewCategoryRepository(mt.DB)

		id := uuid.NewString()
		now := time.Now().UTC().Truncate(time.Millisecond)

		pagination := common.Pagination{
			Page: 1,
			Size: 1,
		}

		mt.AddMockResponses([]primitive.D{
			mtest.CreateCursorResponse(1,
				"foo.bar",
				mtest.FirstBatch,
				bson.D{
					{Key: "n", Value: 1},
				}),
			mtest.CreateCursorResponse(1,
				"foo.bar",
				mtest.FirstBatch,
				getPrimitiveCategory(id, now)),
			mtest.CreateCursorResponse(0,
				"foo.bar",
				mtest.NextBatch),
		}...)

		var expectedCount int64 = 1
		var expectedLength int = 1

		// Act
		count, resp, err := repo.GetAll(context.Background(), pagination)

		// Assert
		assert.NoError(mt, err)
		assert.Equal(mt, expectedCount, count)
		assert.NotNil(mt, resp)
		assert.Len(mt, resp, expectedLength)
	})
}

func TestCategoryRepository_Update(t *testing.T) {
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))

	mt.Run("Should update a category", func(mt *mtest.T) {
		// Arrange
		repo := NewCategoryRepository(mt.DB)

		id := uuid.NewString()
		now := time.Now().UTC().Truncate(time.Millisecond)

		mt.AddMockResponses(bson.D{
			{Key: "ok", Value: 1},
			{Key: "value", Value: bson.D{
				{Key: "uuid", Value: id},
				{Key: "title", Value: "category title"},
				{Key: "description", Value: "category description"},
				{Key: "created_at", Value: now},
				{Key: "updated_at", Value: now},
			}},
		})

		category := entity.NewCategory("category title", "category description", now)

		// Act
		err := repo.Update(context.Background(), category)

		// Assert
		assert.NoError(mt, err)
	})

	mt.Run("Should return error when Decode fails", func(mt *mtest.T) {
		// Arrange
		repo := NewCategoryRepository(mt.DB)

		now := time.Now().UTC().Truncate(time.Millisecond)

		mt.AddMockResponses(mongoInvalidOperation)

		category := entity.NewCategory("category title", "category description", now)

		// Act
		err := repo.Update(context.Background(), category)

		// Assert
		assert.Error(mt, err)
	})

	mt.Run("Should return error when Id is not found", func(mt *mtest.T) {
		// Arrange
		repo := NewCategoryRepository(mt.DB)

		now := time.Now().UTC().Truncate(time.Millisecond)

		mt.AddMockResponses(mtest.CreateCursorResponse(0, "foo.bar", mtest.FirstBatch))

		category := entity.NewCategory("category title", "category description", now)

		// Act
		err := repo.Update(context.Background(), category)

		// Assert
		assert.Error(mt, err)
		assert.ErrorIs(mt, err, ErrCategoryNotFound)
	})
}

func TestCategoryRepository_Delete(t *testing.T) {
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))

	mt.Run("Should delete a category", func(mt *mtest.T) {
		// Arrange
		repo := NewCategoryRepository(mt.DB)

		id := uuid.NewString()

		mt.AddMockResponses(bson.D{
			{Key: "ok", Value: 1},
			{Key: "acknowledged", Value: true}, // This indicates that deletion was okay
			{Key: "n", Value: 1},               // This represents how many documents were deleted
		})

		// Act
		err := repo.Delete(context.Background(), id)

		// Assert
		assert.NoError(mt, err)
	})

	mt.Run("Should return error when DeleteOne fails", func(mt *mtest.T) {
		// Arrange
		repo := NewCategoryRepository(mt.DB)

		id := uuid.NewString()

		mt.AddMockResponses(mongoInvalidOperation)

		// Act
		err := repo.Delete(context.Background(), id)

		// Assert
		assert.Error(mt, err)
	})

	mt.Run("Should return error when nothing was deleted", func(mt *mtest.T) {
		// Arrange
		repo := NewCategoryRepository(mt.DB)

		id := uuid.NewString()

		mt.AddMockResponses(bson.D{
			{Key: "ok", Value: 1},
			{Key: "acknowledged", Value: true}, // This indicates that deletion was okay
			{Key: "n", Value: 0},               // This represents how many documents were deleted
		})

		// Act
		err := repo.Delete(context.Background(), id)

		// Assert
		assert.Error(mt, err)
	})
}

func TestCategoryRepository_getOneByField(t *testing.T) {
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))

	mt.Run("Should get a category by ID", func(mt *mtest.T) {
		// Arrange
		repo := NewCategoryRepository(mt.DB)

		id := uuid.NewString()
		now := time.Now().UTC().Truncate(time.Millisecond)

		mt.AddMockResponses(mtest.CreateCursorResponse(1,
			"foo.bar",
			mtest.FirstBatch,
			getPrimitiveCategory(id, now)))

		expected := entity.Category{
			UUID:        id,
			Title:       "category title",
			Description: "category description",
			CreatedAt:   now,
			UpdatedAt:   now,
		}

		// Act
		resp, err := repo.getOneByField(context.Background(), "uuid", id)

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
		resp, err := repo.getOneByField(context.Background(), "uuid", id)

		// Assert
		assert.Error(mt, err)
		assert.Empty(mt, resp)
	})

	mt.Run("Should return error when nothing is found", func(mt *mtest.T) {
		// Arrange
		repo := NewCategoryRepository(mt.DB)

		id := uuid.NewString()

		mt.AddMockResponses(mtest.CreateCursorResponse(0,
			"foo.bar",
			mtest.FirstBatch))

		// Act
		resp, err := repo.getOneByField(context.Background(), "uuid", id)

		// Assert
		assert.Error(mt, err)
		assert.ErrorIs(mt, err, ErrCategoryNotFound)
		assert.Empty(mt, resp)
	})
}

func TestCategoryRepository_getManyByFieldPaginated(t *testing.T) {
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))

	mt.Run("Should return a list of categories paginated", func(mt *mtest.T) {
		// Arrange
		repo := NewCategoryRepository(mt.DB)

		id := uuid.NewString()
		now := time.Now().UTC().Truncate(time.Millisecond)

		pagination := common.Pagination{
			Page: 1,
			Size: 1,
		}

		mt.AddMockResponses([]primitive.D{
			mtest.CreateCursorResponse(1,
				"foo.bar",
				mtest.FirstBatch,
				bson.D{
					{Key: "n", Value: 1},
				}),
			mtest.CreateCursorResponse(1,
				"foo.bar",
				mtest.FirstBatch,
				getPrimitiveCategory(id, now)),
			mtest.CreateCursorResponse(0,
				"foo.bar",
				mtest.NextBatch),
		}...)

		var expectedCount int64 = 1
		var expectedLength int = 1

		// Act
		count, resp, err := repo.getManyByFieldPaginated(context.Background(), bson.M{"uuid": id}, pagination)

		// Assert
		assert.NoError(mt, err)
		assert.Equal(mt, expectedCount, count)
		assert.NotNil(mt, resp)
		assert.Len(mt, resp, expectedLength)
	})

	mt.Run("Should return error when CountDocuments fails", func(mt *mtest.T) {
		// Arrange
		repo := NewCategoryRepository(mt.DB)

		id := uuid.NewString()

		pagination := common.Pagination{
			Page: 1,
			Size: 1,
		}

		mt.AddMockResponses([]primitive.D{
			mongoInvalidOperation,
		}...)

		var expectedCount int64 = 0
		var expectedLength int = 0

		// Act
		count, resp, err := repo.getManyByFieldPaginated(context.Background(), bson.M{"uuid": id}, pagination)

		// Assert
		assert.Error(mt, err)
		assert.Equal(mt, expectedCount, count)
		assert.Len(mt, resp, expectedLength)
	})

	mt.Run("Should return error when Find fails", func(mt *mtest.T) {
		// Arrange
		repo := NewCategoryRepository(mt.DB)

		id := uuid.NewString()

		pagination := common.Pagination{
			Page: 1,
			Size: 1,
		}

		mt.AddMockResponses([]primitive.D{
			mtest.CreateCursorResponse(1,
				"foo.bar",
				mtest.FirstBatch,
				bson.D{
					{Key: "n", Value: 1},
				}),
			mongoInvalidOperation,
		}...)

		var expectedCount int64 = 1
		var expectedLength int = 0

		// Act
		count, resp, err := repo.getManyByFieldPaginated(context.Background(), bson.M{"uuid": id}, pagination)

		// Assert
		assert.Error(mt, err)
		assert.Equal(mt, expectedCount, count)
		assert.Len(mt, resp, expectedLength)
	})

	mt.Run("Should return error when Decode fails", func(mt *mtest.T) {
		// Arrange
		repo := NewCategoryRepository(mt.DB)

		id := uuid.NewString()

		pagination := common.Pagination{
			Page: 1,
			Size: 1,
		}

		mt.AddMockResponses([]primitive.D{
			mtest.CreateCursorResponse(1,
				"foo.bar",
				mtest.FirstBatch,
				bson.D{
					{Key: "n", Value: 1},
				}),
			mtest.CreateCursorResponse(1,
				"foo.bar",
				mtest.FirstBatch,
				bson.D{
					{Key: "uuid", Value: false},
				}),
		}...)

		var expectedCount int64 = 1
		var expectedLength int = 0

		// Act
		count, resp, err := repo.getManyByFieldPaginated(context.Background(), bson.M{"uuid": id}, pagination)

		// Assert
		assert.Error(mt, err)
		assert.Equal(mt, expectedCount, count)
		assert.Len(mt, resp, expectedLength)
	})

	mt.Run("Should return error from Cursor", func(mt *mtest.T) {
		// Arrange
		repo := NewCategoryRepository(mt.DB)

		id := uuid.NewString()
		now := time.Now().UTC().Truncate(time.Millisecond)

		pagination := common.Pagination{
			Page: 1,
			Size: 1,
		}

		mt.AddMockResponses([]primitive.D{
			mtest.CreateCursorResponse(1,
				"foo.bar",
				mtest.FirstBatch,
				bson.D{
					{Key: "n", Value: 1},
				}),
			mtest.CreateCursorResponse(1,
				"foo.bar",
				mtest.FirstBatch,
				getPrimitiveCategory(id, now)),
			mongoInvalidOperation,
		}...)

		var expectedCount int64 = 1
		var expectedLength int = 0

		// Act
		count, resp, err := repo.getManyByFieldPaginated(context.Background(), bson.M{"uuid": id}, pagination)

		// Assert
		assert.Error(mt, err)
		assert.Equal(mt, expectedCount, count)
		assert.Len(mt, resp, expectedLength)
	})
}
