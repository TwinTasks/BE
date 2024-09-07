package migrations

import (
	dbProjectUsers "twintask/features/projectUsers/repository"
	dbProject "twintask/features/projects/repository"
	dbTask "twintask/features/tasks/repository"
	dbUser "twintask/features/users/repository"

	"gorm.io/gorm"
)

func InitMigration(db *gorm.DB) {
	db.AutoMigrate(&dbUser.Users{})
	db.AutoMigrate(&dbProject.Projects{})
	db.AutoMigrate(&dbTask.Tasks{})
	db.AutoMigrate(&dbProjectUsers.ProjectUsers{})
};