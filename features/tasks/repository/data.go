package repository

import (
	"gorm.io/gorm"
)

type Tasks struct {
	gorm.Model
	ProjectID 	uint
	Title 		string
	Description string
}