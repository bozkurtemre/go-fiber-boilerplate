package repository

import (
	"boilerplate/models"
	"gorm.io/gorm"
)

type UserRepository interface {
	GetAll(limit int) ([]models.User, error)
	GetByID(id uint) (models.User, error)
	Create(user models.User) error
	Update(user models.User) error
	Delete(id uint) error
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db}
}

func (r *userRepository) GetAll(limit int) ([]models.User, error) {
	var users []models.User
	result := r.db.Limit(limit).Find(&users)
	return users, result.Error
}

func (r *userRepository) GetByID(id uint) (models.User, error) {
	var user models.User
	result := r.db.First(&user, id)
	return user, result.Error
}

func (r *userRepository) Create(user models.User) error {
	result := r.db.Create(&user)
	return result.Error
}

func (r *userRepository) Update(user models.User) error {
	result := r.db.Save(&user)
	return result.Error
}

func (r *userRepository) Delete(id uint) error {
	result := r.db.Delete(&models.User{}, id)
	return result.Error
}
