package category

import (
	"context"
	"errors"

	"github.com/jfelipearaujo-org/ms-product-catalog/internal/entity"
	"github.com/jfelipearaujo-org/ms-product-catalog/internal/repository"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	CategoryCollection = "category"
)

var (
	ErrCategoryNotFound = errors.New("category not found")
)

type CategoryRepository struct {
	collection *mongo.Collection
}

func NewCategoryRepository(db *mongo.Database) *CategoryRepository {
	return &CategoryRepository{
		collection: db.Collection(CategoryCollection),
	}
}

func (repo *CategoryRepository) Create(ctx context.Context, category *entity.Category) error {
	_, err := repo.collection.InsertOne(ctx, category)

	return err
}

func (repo *CategoryRepository) GetByID(ctx context.Context, id string) (*entity.Category, error) {
	resp := repo.collection.FindOne(ctx, bson.M{"uuid": id})

	var category entity.Category

	if err := resp.Decode(&category); err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, ErrCategoryNotFound
		}

		return nil, err
	}

	return &category, nil
}

func (repo *CategoryRepository) GetByTitle(ctx context.Context, title string) (*entity.Category, error) {
	resp := repo.collection.FindOne(ctx, bson.M{"title": title})

	var category entity.Category

	if err := resp.Decode(&category); err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, ErrCategoryNotFound
		}

		return nil, err
	}

	return &category, nil
}

func (repo *CategoryRepository) GetAll(ctx context.Context, filter repository.Pagination) (int64, []entity.Category, error) {
	var categories []entity.Category

	countOpts := options.Count().SetHint("_id_")
	count, err := repo.collection.CountDocuments(ctx, bson.D{}, countOpts)
	if err != nil {
		return 0, categories, err
	}

	skip := filter.Page*filter.Size - filter.Size

	findOpts := options.Find().
		SetLimit(filter.Size).
		SetSkip(skip).
		SetSort(bson.D{{
			Key:   "created_at",
			Value: 1,
		}})

	cursor, err := repo.collection.Find(ctx, bson.D{}, findOpts)
	if err != nil {
		return 0, categories, err
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		category := entity.Category{}
		if err := cursor.Decode(&category); err != nil {
			return 0, categories, err
		}

		categories = append(categories, category)
	}

	if err := cursor.Err(); err != nil {
		return count, categories, err
	}

	return count, categories, nil
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
			return ErrCategoryNotFound
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
		return ErrCategoryNotFound
	}

	return err
}
