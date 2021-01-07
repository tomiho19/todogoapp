package models

import (
	"time"
)

type Todo struct {
	ID int `schema:"-"`
	Body string `schema:"body"`
	CreatedAt time.Time `schema:"-"`
	UpdatedAt time.Time `schema:"-"`
}