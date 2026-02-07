package repository

import "github.com/claytonssmint/task-manager-go/internal/model"

type TaskRepository interface {
	Create(task *model.Task) error
	FindAll() ([]model.Task, error)
	FindByID(id int64) (*model.Task, error)
	Update(task *model.Task) error
	Delete(id int64) error
}
