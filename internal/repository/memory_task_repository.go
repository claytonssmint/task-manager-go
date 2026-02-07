package repository

import (
	"errors"
	"sync"
	"time"

	"github.com/claytonssmint/task-manager-go/internal/model"
)

type MemoryTaskRepository struct {
	mu     sync.Mutex
	tasks  map[int64]model.Task
	nextID int64
}

func NewMemoryTaskRepository() *MemoryTaskRepository {
	return &MemoryTaskRepository{
		tasks:  make(map[int64]model.Task),
		nextID: 1,
	}
}

// Create
func (r *MemoryTaskRepository) Create(task *model.Task) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	task.ID = r.nextID
	task.CreatedAt = time.Now()
	task.UpdatedAt = time.Now()

	r.tasks[task.ID] = *task
	r.nextID++

	return nil
}

// FindAll
func (r *MemoryTaskRepository) FindAll() ([]model.Task, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	var list []model.Task

	for _, task := range r.tasks {
		list = append(list, task)
	}

	return list, nil
}

// FindByID
func (r *MemoryTaskRepository) FindByID(id int64) (*model.Task, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	task, ok := r.tasks[id]
	if !ok {
		return nil, errors.New("task not found")
	}

	return &task, nil
}

// Update
func (r *MemoryTaskRepository) Update(task *model.Task) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	_, ok := r.tasks[task.ID]
	if !ok {
		return errors.New("task not found")
	}

	task.UpdatedAt = time.Now()
	r.tasks[task.ID] = *task

	return nil
}

// Delete
func (r *MemoryTaskRepository) Delete(id int64) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, ok := r.tasks[id]; !ok {
		return errors.New("task not found")
	}

	delete(r.tasks, id)

	return nil
}
