package entity

import (
	"app/graph/model"
	"time"
)

type Todo struct {
	ID          uint64     `json:"id" gorm:"primary_key"`
	Title       string     `json:"title"`
	Description string     `json:"description"`
	UserId      uint64     `json:"userId"`
	User        User       `json:"-"`
	CreatedAt   time.Time  `json:"createdAt"`
	UpdatedAt   time.Time  `json:"updatedAt"`
	DeletedAt   *time.Time `json:"deletedAt"`
}

func ToModelTodo(t *Todo) *model.Todo {
	return &model.Todo{
		ID:          string(rune(t.ID)),
		Title:       t.Title,
		Description: t.Description,
	}
}

func ToModelTodos(ts []*Todo) []*model.Todo {
	var todos []*model.Todo
	for i := 0; i < len(ts); i++ {
		todos = append(todos, ToModelTodo(ts[i]))
	}

	return todos
}
