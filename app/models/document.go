package models

type Document struct {
	Id      int    `form:"-" json:"-" db:"id"`
	Content string `form:"content" json:"content" db:"content"`
}
