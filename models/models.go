package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID        		uuid.UUID `gorm:"primarykey"`

	Todo 			[]Todo
}

type Todo struct {
	gorm.Model
	ID        		uuid.UUID `gorm:"primarykey"`
	Content			string
	Description		string

	UserId			uuid.UUID
}