package types

import (
	"github.com/chenminhua/data_generator/constant"
	"time"
)

type User struct {
	Id        int64 `gorm:"primaryKey;autoIncrement"`
	Name      string
	Email     string
	Password  string
	Bio       string
	Avatar    string
	CreatedAt time.Time `gorm:"autoCreateTime:milli"`
	UpdatedAt time.Time `gorm:"autoUpdateTime:milli"`
}

func NewUser(idx int) *User {
	username := constant.GetName(idx)
	return &User{
		Name:     username,
		Email:    username + "@gmail.com",
		Password: "1234wqqwqer",
		Bio:      "",
		Avatar:   "",
	}
}

type Article struct {
	Id        int64 `gorm:"primaryKey;autoIncrement"`
	Title     string
	Body      string
	Author    Author
	CreatedAt time.Time `gorm:"autoCreateTime:milli"`
	UpdatedAt time.Time `gorm:"autoUpdateTime:milli"`
}

type Author struct {
	Id   int64
	Name string
}