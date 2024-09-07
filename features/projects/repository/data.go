package repository

import (
	dbProjectUsers "twintask/features/projectUsers/repository"
	"twintask/features/tasks/repository"

	"gorm.io/gorm"
)

type Projects struct {
	gorm.Model
	UserID 		 uint
	Title 		 string
	Description  string
	Tasks 		 repository.Tasks         `gorm:"foreignKey:ProjectID"`
	ProjectUsers dbProjectUsers.ProjectUsers  `gorm:"foreignKey:ProjectID"`
}