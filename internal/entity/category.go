package entity

import (
	"time"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Category struct {
	Id   primitive.ObjectID `bson:"_id,omitempty" json:"-"`
	UUID string             `bson:"uuid" json:"id"`

	Title       string `bson:"title" json:"title"`
	Description string `bson:"description" json:"description"`

	CreatedAt time.Time `bson:"created_at" json:"created_at"`
	UpdatedAt time.Time `bson:"updated_at" json:"updated_at"`
}

func NewCategory(
	title string,
	description string,
	now time.Time) *Category {
	return &Category{
		UUID:        uuid.New().String(),
		Title:       title,
		Description: description,
		CreatedAt:   now,
		UpdatedAt:   now,
	}
}

func (c *Category) Update(
	title string,
	description string,
	now time.Time) {
	c.Title = title
	c.Description = description
	c.UpdatedAt = now
}
