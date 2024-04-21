package entity

import (
	"time"

	"github.com/google/uuid"
)

type Category struct {
	UUID string `bson:"uuid"`

	Title       string `bson:"title"`
	Description string `bson:"description"`

	CreatedAt time.Time `bson:"created_at"`
	UpdatedAt time.Time `bson:"updated_at"`
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
