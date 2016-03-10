package models

import (
	"time"
)

type User struct {
	ID        int32
	Name      string
	Email     string
	Password  string
	Desc      string
	CreatedAt time.Time
	LastLogin time.Time
}

type Article struct {
	ID        uint64
	Title     string
	Catgory   string
	Keywords  string
	Content   string
	CreatedAt time.Time
	UpdatedAt time.Time
	CreatedBy int32
	Hits      uint64
}
