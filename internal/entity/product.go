package entity

import (
	"time"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Product struct {
	Id   primitive.ObjectID `bson:"_id,omitempty" json:"-"`
	UUID string             `bson:"uuid" json:"id"`

	Title       string   `bson:"title" json:"title"`
	Description string   `bson:"description" json:"description"`
	Price       float64  `bson:"price" json:"price"`
	Category    Category `bson:"category" json:"category"`

	CreatedAt time.Time `bson:"created_at" json:"created_at"`
	UpdatedAt time.Time `bson:"updated_at" json:"updated_at"`
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
