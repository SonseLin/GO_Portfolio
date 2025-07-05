package storage

import (
	"rtdapp/internal/errors"

	"github.com/google/uuid"
)

type Task struct {
	ID          uint32 `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type Storage struct {
	tasks    []Task
	capacity int
	len      int
	TaskMap  map[string]Task
}

func NewTask(name, description string) *Task {
	return &Task{
		ID:          uuid.New().ID(),
		Name:        name,
		Description: description,
	}
}

func NewStorage(capacity int) *Storage {
	return &Storage{
		tasks:    make([]Task, 0, capacity),
		capacity: capacity,
		len:      0,
		TaskMap:  make(map[string]Task),
	}
}

func (s *Storage) AddTask(task Task) error {
	if s.len < s.capacity {
		if _, ok := s.TaskMap[task.Name]; !ok {
			s.tasks = append(s.tasks, task)
			s.len++
			s.TaskMap[task.Name] = task
			return nil
		}
		return errors.TaskAlreadyExists
	}
	return errors.StorageIsFull
}

func (s *Storage) GetTask(name string) (*Task, error) {
	if task, ok := s.TaskMap[name]; ok {
		return &task, nil
	}
	return &Task{}, errors.TaskNotFound
}

func (s *Storage) DeleteTask(name string) error {
	if _, ok := s.TaskMap[name]; ok {
		delete(s.TaskMap, name)
		s.len--
		return nil
	}
	return errors.TaskNotFound
}
