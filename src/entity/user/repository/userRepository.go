package userRepository

import (
	"phone-contact/infrastructure/repository"

	"gorm.io/gorm"
)

type UserRepository struct {
	Repo *repository.Repository
}

func CreateUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{
		Repo: repository.CreateRepository(db),
	}
}
