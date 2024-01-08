package models

import "time"
//Article article info model.
type Article struct {
	Model
	Title   string `form:"title"`
	Content string `form:"content"`
}
//Model base model.
type Model struct {
	ID        uint64     `form:"id"`
	CreatedAt time.Time  `binding:"-" form:"-"`
	UpdatedAt time.Time  `binding:"-" form:"-"`
	DeletedAt *time.Time `binding:"-" form:"-"`
}
