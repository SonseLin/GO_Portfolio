package errors

import "errors"

var (
	TaskNotFound      = errors.New("task not found")
	TaskAlreadyExists = errors.New("task already exists")
	StorageIsFull     = errors.New("storage full")
)
