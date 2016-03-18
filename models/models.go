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

type ArtList struct {
	ID            uint64
	CreatedBy     int32
	CreatedByName string
	Title         string
	LastReply     time.Time
	LastReplyBy   int32
	LastReplyName string
}

func UserQuery(u *User) (err error) {
    return nil
}

func ArtQuery(art *Article) (err error) {
    return nil
}

func ArtListQuery(page int, sort int) (list []ArtList, err error) {
    return nil, nil
}

func ArtContentQuery(key string) (list []ArtList, err error) {
    return nil, nil
}