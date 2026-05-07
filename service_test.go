package main

import (
	"testing"
)

type mockRepo struct{ tasks []Task }

func (m *mockRepo) GetAll() ([]Task, error)   { return m.tasks, nil }
func (m *mockRepo) Save(t []Task) error        { m.tasks = t; return nil }

func TestAdd(t *testing.T) {
	svc := NewTaskService(&mockRepo{})
	if err := svc.Add("Buy groceries", Low); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	tasks, _ := svc.List()
	if len(tasks) != 1 {
		t.Errorf("expected 1 task, got %d", len(tasks))
	}
}

func TestComplete(t *testing.T) {
	repo := &mockRepo{}
	svc  := NewTaskService(repo)
	_ = svc.Add("Write report", High)
	if err := svc.Complete(1); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	tasks, _ := svc.repo.GetAll()
	if !tasks[0].Done {
		t.Error("task should be marked done")
	}
}

func TestDelete(t *testing.T) {
	repo := &mockRepo{}
	svc  := NewTaskService(repo)
	_ = svc.Add("Temp task", Low)
	if err := svc.Delete(1); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	tasks, _ := svc.repo.GetAll()
	if len(tasks) != 0 {
		t.Errorf("expected 0 tasks after delete, got %d", len(tasks))
	}
}

func TestListSortedByPriority(t *testing.T) {
	repo := &mockRepo{}
	svc  := NewTaskService(repo)
	_ = svc.Add("Low task",    Low)
	_ = svc.Add("High task",   High)
	_ = svc.Add("Medium task", Medium)
	tasks, _ := svc.List()
	if tasks[0].Priority != High {
		t.Error("first task should be High priority")
	}
}