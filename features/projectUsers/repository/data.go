package repository

import "gorm.io/gorm"

type ProjectUsers struct {
	gorm.Model
	ProjectID uint
	UserID	  uint
}