package model

import (
	"time"
)

type TodoList struct {
	ID          string    `json:"id"`
	Todo        string    `json:"todo" validate:"required,omitempty,min=1"`
	Iscompleted bool      `json:"is_completed"`
	CreatedAt   time.Time `json:"created_at"`
}

type TodoListUpdate struct {
	Todo string `json:"todo" validate:"omitempty,min=1"`
	IsCompleted *bool  `json:"is_completed" validate:"required"`
}
