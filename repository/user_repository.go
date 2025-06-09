package repository

import (
	"church_user/domain"
)

type UserRepository interface {
	Save(user *domain.User) error
	FindByID(id uint) (*domain.User, error)
	FindAll() ([]domain.User, error)
	Update(user *domain.User) error
	Delete(id uint) error
}
