package category

import (
	"context"
	"errors"

	"github.com/jfelipearaujo-org/ms-product-catalog/internal/common"
	"github.com/jfelipearaujo-org/ms-product-catalog/internal/entity"
	"github.com/jfelipearaujo-org/ms-product-catalog/internal/repository"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	CategoryCollection = "category"
)

type CategoryRepository struct {
	collection *mongo.Collection
}

func NewRepository(db *mongo.Database) *CategoryRepository {
	return &CategoryRepository{
		collection: db.Collection(CategoryCollection),
	}
}

func (repo *CategoryRepository) Create(ctx context.Context, category *entity.Category) error {
	res, err := repo.collection.InsertOne(ctx, category)
	if err != nil {
		return err
	}

	category.Id = res.InsertedID.(primitive.ObjectID)

	return err
}

func (repo *CategoryRepository) GetByID(ctx context.Context, id string) (entity.Category, error) {
	return repo.getOneByField(ctx, "uuid", id)
}

func (repo *CategoryRepository) GetByTitle(ctx context.Context, title string) (entity.Category, error) {
	return repo.getOneByField(ctx, "title", title)
}

func (repo *CategoryRepository) GetAll(ctx context.Context, pagination common.Pagination, filter repository.GetAllCategoriesFilter) (int64, []entity.Category, error) {
	filters := []bson.M{}

	if filter.Title != "" {
		filters = append(filters, bson.M{"title": filter.Title})
	}

	var query interface{}

	if len(filters) > 0 {
		query = bson.M{"$and": filters}
	} else {
		query = bson.M{}
	}

	return repo.getManyByFieldPaginated(ctx, query, pagination)
}

func (repo *CategoryRepository) Update(ctx context.Context, category *entity.Category) error {
	query := bson.M{"$set": bson.M{
		"title":       category.Title,
		"description": category.Description,
		"updated_at":  category.UpdatedAt,
	}}

	if err := repo.collection.FindOneAndUpdate(
		ctx,
		bson.M{"uuid": category.UUID},
		query,
		options.FindOneAndUpdate().SetReturnDocument(options.After)).Decode(category); err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return repository.ErrCategoryNotFound
		}

		return err
	}

	return nil
}

func (repo *CategoryRepository) Delete(ctx context.Context, id string) error {
	resp, err := repo.collection.DeleteOne(ctx, bson.M{"uuid": id})
	if err != nil {
		return err
	}

	if resp.DeletedCount == 0 {
		return repository.ErrCategoryNotFound
	}

	return err
}

// private methods
func (repo *CategoryRepository) getOneByField(ctx context.Context, field string, value string) (entity.Category, error) {
	resp := repo.collection.FindOne(ctx, bson.M{field: value})

	var category entity.Category

	if err := resp.Decode(&category); err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return category, repository.ErrCategoryNotFound
		}

		return category, err
	}

	return category, nil
}

func (repo *CategoryRepository) getManyByFieldPaginated(ctx context.Context, query interface{}, pagination common.Pagination) (int64, []entity.Category, error) {
	var categories []entity.Category

	indexModel := mongo.IndexModel{
		Keys: bson.D{{
			Key:   "created_at",
			Value: 1,
		}},
	}

	_, err := repo.collection.Indexes().CreateOne(ctx, indexModel)
	if err != nil {
		return 0, categories, err
	}

	countOpts := options.Count()
	count, err := repo.collection.CountDocuments(ctx, query, countOpts)
	if err != nil {
		return 0, categories, err
	}

	skip := pagination.Page*pagination.Size - pagination.Size

	findOpts := options.Find().
		SetLimit(pagination.Size).
		SetSkip(skip).
		SetSort(bson.D{{
			Key:   "created_at",
			Value: 1,
		}})

	resp, err := repo.collection.Find(ctx, query, findOpts)
	if err != nil {
		return count, categories, err
	}

	if err := resp.All(ctx, &categories); err != nil {
		return count, categories, err
	}

	return count, categories, nil
}
