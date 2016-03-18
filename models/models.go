package models

import (
	"time"
)

var (
    pageSize = 20
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

type ExtArt struct {
	ID            uint64
	CreatedBy     int32
	CreatedByName string
	Title         string
	LastReply     time.Time
	LastReplyBy   int32
	LastReplyName string
}

//query by pk:ID
func UserQuery(u *User) (err error) {
    return DB.First(&u).Error
}

//query by pk:ID
func ArtQuery(art *Article) (err error) {
    return DB.First(&art).Error
}


func ArtListQuery(page int, sort int) (list []ExtArt, err error) {
    err = DB.Table("article").Select("article.id, article.Title, article.category, article.create_at").Joins("inner join user on user.id = article.create_by").Scan(&list).Offset((page - 1) * pageSize).Limit(pageSize).Error
    if err != nil {
        return nil, err
    }
    return list, nil
}

func ArtContentQuery(key string) (list []ExtArt, err error) {
    return nil, nil
}