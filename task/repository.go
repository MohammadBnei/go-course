package task

import (
	"errors"

	"gorm.io/gorm"
)

type Repository interface {
	Store(task Task) (Task, error)
	FetchAll() ([]Task, error)
	FetchById(id int) (Task, error)
	Update(id int, inputTask InputTask) (Task, error)
	Delete(id int) error
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) Store(task Task) (Task, error) {
	err := r.db.Create(&task).Error
	if err != nil {
		return task, err
	}

	return task, nil
}

func (r *repository) FetchAll() ([]Task, error) {
	var tasks []Task
	err := r.db.Find(&tasks).Error
	if err != nil {
		return tasks, err
	}

	return tasks, nil
}

func (r *repository) FetchById(id int) (Task, error) {
	var task Task

	err := r.db.Where(&Task{ID: id}).First(&task).Error
	if err != nil {
		return task, err
	}

	return task, nil
}

func (r *repository) Update(id int, inputTask InputTask) (Task, error) {
	task, err := r.FetchById(id)
	if err != nil {
		return task, err
	}

	task.Name = inputTask.Name
	task.Description = inputTask.Description

	err = r.db.Save(&task).Error
	if err != nil {
		return task, err
	}

	return task, nil
}

func (r *repository) Delete(id int) error {
	task := &Task{ID: id}
	tx := r.db.Delete(task)
	if tx.Error != nil {
		return tx.Error
	}

	if tx.RowsAffected == 0 {
		return errors.New("Task not found")
	}

	return nil
}
