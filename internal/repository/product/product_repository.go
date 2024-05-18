package product

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
	ProductCollection = "product"
)

type ProductRepository struct {
	collection *mongo.Collection
}

func NewRepository(db *mongo.Database) *ProductRepository {
	return &ProductRepository{
		collection: db.Collection(ProductCollection),
	}
}

func (repo *ProductRepository) Create(ctx context.Context, product *entity.Product) error {
	res, err := repo.collection.InsertOne(ctx, product)
	if err != nil {
		return err
	}

	product.Id = res.InsertedID.(primitive.ObjectID)

	return err
}

func (repo *ProductRepository) GetByID(ctx context.Context, id string) (entity.Product, error) {
	return repo.getOneByField(ctx, "uuid", id)
}

func (repo *ProductRepository) GetByTitle(ctx context.Context, title string) (entity.Product, error) {
	return repo.getOneByField(ctx, "title", title)
}

func (repo *ProductRepository) GetByCategoryID(ctx context.Context, categoryId string, pagination common.Pagination) (int64, []entity.Product, error) {
	return repo.getManyByFieldPaginated(ctx, bson.M{"category.uuid": categoryId}, pagination)
}

func (repo *ProductRepository) GetByCategoryTitle(ctx context.Context, categoryTitle string, pagination common.Pagination) (int64, []entity.Product, error) {
	return repo.getManyByFieldPaginated(ctx, bson.M{"category.title": categoryTitle}, pagination)
}

func (repo *ProductRepository) GetAll(ctx context.Context, pagination common.Pagination, filter repository.GetAllProductsFilter) (int64, []entity.Product, error) {
	filters := []bson.M{}

	if filter.Title != "" {
		filters = append(filters, bson.M{"title": filter.Title})
	}

	if filter.CategoryTitle != "" {
		filters = append(filters, bson.M{"category.title": filter.CategoryTitle})
	}

	var query interface{}

	if len(filters) > 0 {
		query = bson.M{"$and": filters}
	} else {
		query = bson.M{}
	}

	return repo.getManyByFieldPaginated(ctx, query, pagination)
}

func (repo *ProductRepository) Update(ctx context.Context, product *entity.Product) error {
	query := bson.M{"$set": bson.M{
		"title":       product.Title,
		"description": product.Description,
		"price":       product.Price,
		"category":    product.Category,
		"updated_at":  product.UpdatedAt,
	}}

	if err := repo.collection.FindOneAndUpdate(
		ctx,
		bson.M{"uuid": product.UUID},
		query,
		options.FindOneAndUpdate().SetReturnDocument(options.After)).Decode(product); err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return repository.ErrProductNotFound
		}

		return err
	}

	return nil
}

func (repo *ProductRepository) Delete(ctx context.Context, id string) error {
	resp, err := repo.collection.DeleteOne(ctx, bson.M{"uuid": id})
	if err != nil {
		return err
	}

	if resp.DeletedCount == 0 {
		return repository.ErrProductNotFound
	}

	return err
}

// private methods
func (repo *ProductRepository) getOneByField(ctx context.Context, field string, value string) (entity.Product, error) {
	resp := repo.collection.FindOne(ctx, bson.M{field: value})

	var product entity.Product

	if err := resp.Decode(&product); err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return product, repository.ErrProductNotFound
		}

		return product, err
	}

	return product, nil
}

func (repo *ProductRepository) getManyByFieldPaginated(ctx context.Context, query interface{}, pagination common.Pagination) (int64, []entity.Product, error) {
	var products []entity.Product

	indexModel := mongo.IndexModel{
		Keys: bson.D{{
			Key:   "created_at",
			Value: 1,
		}},
	}

	_, err := repo.collection.Indexes().CreateOne(ctx, indexModel)
	if err != nil {
		return 0, products, err
	}

	countOpts := options.Count()
	count, err := repo.collection.CountDocuments(ctx, query, countOpts)
	if err != nil {
		return 0, products, err
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
		return count, products, err
	}

	if err := resp.All(ctx, &products); err != nil {
		return count, products, err
	}

	return count, products, nil
}
