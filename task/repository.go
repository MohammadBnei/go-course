package task

import "gorm.io/gorm"

type Repository interface {
	Insert(task Task) (Task, error)
	ListAll() ([]Task, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) Insert(task Task) (Task, error) {
	err := r.db.Create(&task).Error
	if err != nil {
		return task, err
	}

	return task, nil
}

func (r *repository) ListAll() ([]Task, error) {
	var tasks []Task
	err := r.db.Find(&tasks).Error
	if err != nil {
		return nil, err
	}

	return tasks, nil
}
