package models

import (
	"time"
    "eden/config"
)

var (
    pageSize = 20
)

type Note struct {
    ID        int
    Title     string
    Content   string
    CreatedBy int
    CreatedAt time.Time
    UpdatedAt time.Time
}

type User struct {
	ID        int
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
	CreatedBy int
	Hits      uint64
}

type ExtArt struct {
	ID            uint64
	CreatedBy     int
	CreatedByName string
	Title         string
	LastReply     time.Time
	LastReplyBy   int
	LastReplyName string
}

func (u User) IsAdmin() bool {
    return u.Name == config.Config.Admin
}

//user query by pk:ID
func UserQuery(u *User) (err error) {
    return DB.First(&u).Error
}

func UserExsist(u *User) bool {
    return !DB.First(u).RecordNotFound()
}

//article query by pk:ID
func ArtQuery(art *Article) (err error) {
    return DB.First(&art).Error
}

//show article
func ShowArt(id uint64) (exArt *ExtArt, err error) {
    err = DB.Table("article").Select("article.id, article.Title, article.category, article.create_at, user.name").Joins("inner join user on user.id = article.create_by").Scan(&exArt).Error
    if err != nil {
        return nil, err
    }
    return exArt, nil
}

//note query by pk:ID
func NoteQuery(n *Note) (err error) {
    return DB.First(&n).Error
}

//update multifield
func UpdateArt(art *Article) error {
    return DB.Update("title", art.Title).Update("content", art.Content).Error
}

//delete article
func DeleteArt(id uint64) (err error) {
    art := Article{ID: id}
    return DB.Delete(&art).Error
}

func ArtListQuery(page int, sort int) (list []ExtArt, err error) {
    err = DB.Table("article").Select("article.id, article.Title, article.category, article.create_at, user.name").Joins("inner join user on user.id = article.create_by").Scan(&list).Offset((page - 1) * pageSize).Limit(pageSize).Error
    if err != nil {
        return nil, err
    }
    return list, nil
}

func ArtContentQuery(key string) (list []ExtArt, err error) {
    return nil, nil
}

func UpdateUser(user *User) (err error) {
    return DB.Save(user).Error
}

func CreateUser(user *User) (err error) {
    return DB.Create(user).Error
}

func CreateArticle(art *Article) (err error) {
    return DB.Create(art).Error
}