package entity

import (
	"app/graph/model"
	"strconv"
	"time"
)

type User struct {
	ID           uint64     `json:"id" gorm:"primary_key"`
	Name         string     `json:"name"`
	Email        string     `json:"email"`
	HashPassword string     `json:"-"`
	CreatedAt    time.Time  `json:"createdAt"`
	UpdatedAt    time.Time  `json:"updatedAt"`
	DeletedAt    *time.Time `json:"deletedAt"`
}

func ToModelUser(u *User) *model.User {
	return &model.User{
		ID:   strconv.Itoa(int(u.ID)),
		Name: u.Name,
	}
}
