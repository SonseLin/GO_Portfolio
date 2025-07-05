package internal

import "errors"

type Task struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type Storage struct {
	tasks    []Task
	capacity int
	len      int
	TaskMap  map[string]Task
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
		return errors.New("Task already exists")
	}
	return errors.New("Storage is full")
}

func (s *Storage) GetTask(name string) (Task, error) {
	if task, ok := s.TaskMap[name]; ok {
		return task, nil
	}
	return Task{}, errors.New("Task not found")
}

func (s *Storage) DeleteTask(name string) error {
	if _, ok := s.TaskMap[name]; ok {
		delete(s.TaskMap, name)
		s.len--
		return nil
	}
	return errors.New("Task not found")
}
