package storage

import (
	"rtdapp/internal/errors"
	"testing"
)

func TestCreateStorage(t *testing.T) {
	cap := 10
	stor := NewStorage(cap)
	if stor == nil {
		t.Fatal("Newstorage returned nil")
	}
	if stor.capacity != cap {
		t.Fatalf(
			"Newstorage returned wrong capacity. Expected %d, got %d",
			cap, stor.capacity,
		)
	}
	if stor.TaskMap == nil {
		t.Fatal(
			"Newstorage returned nil TaskMap",
		)
	}
}

func TestAddTaskBeyondCapacity(t *testing.T) {
	stor := NewStorage(1)
	task := NewTask("tmp", "temporal task")
	task_2 := NewTask("tmp2", "temporal task")
	stor.AddTask(*task)
	err := stor.AddTask(*task_2)
	if err != errors.StorageIsFull {
		t.Fatal(
			"Storage is full, but AddTask returned no error",
		)
	}
}

func TestTaskAlreadyExists(t *testing.T) {
	stor := NewStorage(2)
	task := NewTask("tmp", "temporal task")
	stor.AddTask(*task)
	err := stor.AddTask(*task)
	if err != errors.TaskAlreadyExists {
		t.Fatal(
			"Task already exists, but AddTask returned no error",
		)
	}
}

func TestGetTask(t *testing.T) {
	stor := NewStorage(2)
	task := NewTask("tmp", "temporal task")
	stor.AddTask(*task)
	tmp, err := stor.GetTask("tmp")
	if err != nil {
		t.Fatal(
			"GetTask returned error",
		)
	}
	if tmp != task {

	}
}

func TestGetUnexistingTask(t *testing.T) {
	stor := NewStorage(2)
	_, err := stor.GetTask("tmp")
	if err != errors.TaskNotFound {
		t.Fatal(
			"GetTask returned wrong error",
		)
	}
}
