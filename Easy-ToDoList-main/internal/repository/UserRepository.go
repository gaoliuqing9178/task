package repository

import (
	"gorm.io/gorm"
	"todolist/internal/model"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) Create(user *model.User) error {
	return r.db.Create(user).Error
}

func (r *UserRepository) GetByID(id uint) (*model.User, error) {
	var user model.User
	err := r.db.First(&user, id).Error
	return &user, err
}

func (r *UserRepository) Update(user *model.User) error {
	return r.db.Save(user).Error
}

func (r *UserRepository) UpdateByAccount(user *model.User) error {
	tx := r.db.Model(&model.User{}).Where("account = ?", user.Account).Updates(user)
	return tx.Error
}

func (r *UserRepository) Delete(id uint) error {
	return r.db.Delete(&model.User{}, id).Error
}

func (r *UserRepository) GetByAccount(account string) (*model.User, error) {
	var user model.User
	err := r.db.Where("account = ?", account).First(&user).Error
	return &user, err
}
