package entity

import (
	"time"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Product struct {
	Id   primitive.ObjectID `bson:"_id,omitempty"`
	UUID string             `bson:"uuid"`

	Title       string   `bson:"title"`
	Description string   `bson:"description"`
	Price       float64  `bson:"price"`
	Category    Category `bson:"category"`

	CreatedAt time.Time `bson:"created_at"`
	UpdatedAt time.Time `bson:"updated_at"`
}

func NewProduct(
	title string,
	description string,
	price float64,
	category Category,
	now time.Time,
) *Product {
	return &Product{
		UUID:        uuid.New().String(),
		Title:       title,
		Description: description,
		Price:       price,
		Category:    category,
		CreatedAt:   now,
		UpdatedAt:   now,
	}
}

func (p *Product) Update(
	title string,
	description string,
	price float64,
	category Category,
	now time.Time,
) {
	p.Title = title
	p.Description = description
	p.Price = price
	p.Category = category
	p.UpdatedAt = now
}
