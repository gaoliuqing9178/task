package repository

import (
	"gorm.io/gorm"
	"todolist/internal/model"
)

type ToDoRepository struct {
	db *gorm.DB
}

func NewToDoRepository(db *gorm.DB) *ToDoRepository {
	return &ToDoRepository{db: db}
}

func (r *ToDoRepository) Create(todo *model.ToDo) error {
	return r.db.Create(todo).Error
}

func (r *ToDoRepository) GetByID(id uint) (*model.ToDo, error) {
	var todo model.ToDo
	err := r.db.First(&todo, id).Error
	return &todo, err
}

func (r *ToDoRepository) UpdateById(todo *model.ToDo, id uint) error {
	return r.db.Model(&model.ToDo{}).Where("id = ?", id).Updates(todo).Error
}

func (r *ToDoRepository) Delete(id uint) error {
	return r.db.Delete(&model.ToDo{}, id).Error
}

func (r *ToDoRepository) ListByUser(userId uint) ([]model.ToDo, error) {
	var todos []model.ToDo
	err := r.db.Where("user_id = ?", userId).Find(&todos).Error
	return todos, err
}
