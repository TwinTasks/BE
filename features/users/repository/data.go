package repository

import (
	prjUsers "twintask/features/projectUsers/repository"
	dbProject "twintask/features/projects/repository"
	"twintask/features/users"

	"gorm.io/gorm"
)

type Users struct {
	gorm.Model
	Username 	 string
	Email 	  	 string
	Password 	 string
	Projects 	 []dbProject.Projects `gorm:"foreignKey:UserID"`
	ProjectUsers []prjUsers.ProjectUsers `gorm:"foreignKey:UserID"`
}


func ToUserQuery(input users.Users) Users {
	return Users{
		Username: input.Username,
		Email: 	  input.Email,
		Password: input.Password,
	}
}

func (um *Users) ToUserEntity() users.Users{
	return users.Users{
		Username: um.Username,
		Email:    um.Email,
		Password: um.Password,
	}
}