package main

import (
	"time"
)

type CreateTodoRequest struct {
	Title string `json:"title"`
}

type UpdateTodoRequest struct {
	Title  string `json:"title"`
	Status bool   `json:"status"`
}

type Todo struct {
	ID        int       `json:"id"`
	Title     string    `json:"title"`
	Status    bool      `json:"status"`
	CreatedAt time.Time `json:"createdAt"`
}

func NewTodo(title string) *Todo {
	return &Todo{
		Title:     title,
		CreatedAt: time.Now().UTC(),
	}
}
