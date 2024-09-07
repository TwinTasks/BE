package repository

import (
	"twintask/features/users"

	"gorm.io/gorm"
)

type UserModels struct {
	db *gorm.DB
};

func NewUserModels(connect *gorm.DB) users.Query{
	return &UserModels{
		db: connect,
	}
}

// Register 
func (um *UserModels) Register(NewUser users.Users) error {
	cnvData := ToUserQuery(NewUser)
	err := um.db.Create(&cnvData).Error

	if err != nil {
		return err
	}

	return nil;
}

func (um *UserModels) Login(email string) (users.Users, error){
	var result Users;
	err := um.db.Where("email = ?", email).First(&result).Error

	if err != nil {
		return users.Users{}, err
	}

	return result.ToUserEntity(), nil;
}