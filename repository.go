package main

import (
	"encoding/json"
	"os"
)

type Repository interface {
	GetAll() ([]Task, error)
	Save(tasks []Task) error
}

type FileRepository struct{ path string }

func NewFileRepository(path string) *FileRepository {
	return &FileRepository{path: path}
}

func (r *FileRepository) GetAll() ([]Task, error) {
	data, err := os.ReadFile(r.path)
	if os.IsNotExist(err) { return []Task{}, nil }
	if err != nil { return nil, err }
	var tasks []Task
	return tasks, json.Unmarshal(data, &tasks)
}

func (r *FileRepository) Save(tasks []Task) error {
	data, err := json.MarshalIndent(tasks, "", "  ")
	if err != nil { return err }
	return os.WriteFile(r.path, data, 0644)
}