package main

import (
	"errors"
	"sort"
	"time"
)

type TaskService struct{ repo Repository }

func NewTaskService(r Repository) *TaskService { return &TaskService{repo: r} }

func (s *TaskService) Add(title string, p Priority) error {
	tasks, err := s.repo.GetAll()
	if err != nil { return err }
	id := len(tasks) + 1
	tasks = append(tasks, Task{
		ID: id, Title: title,
		Priority: p, CreatedAt: time.Now(),
	})
	return s.repo.Save(tasks)
}

func (s *TaskService) Complete(id int) error {
	tasks, err := s.repo.GetAll()
	if err != nil { return err }
	for i, t := range tasks {
		if t.ID == id { tasks[i].Done = true
			return s.repo.Save(tasks) }
	}
	return errors.New("task not found")
}

func (s *TaskService) Delete(id int) error {
	tasks, err := s.repo.GetAll()
	if err != nil { return err }
	for i, t := range tasks {
		if t.ID == id {
			tasks = append(tasks[:i], tasks[i+1:]...)
			return s.repo.Save(tasks)
		}
	}
	return errors.New("task not found")
}

func (s *TaskService) List() ([]Task, error) {
	tasks, err := s.repo.GetAll()
	if err != nil { return nil, err }
	sort.Slice(tasks, func(i, j int) bool {
		return tasks[i].Priority > tasks[j].Priority
	})
	return tasks, nil
}