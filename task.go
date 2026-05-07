package main

import "time"

type Priority int

const (
	Low Priority = iota
	Medium
	High
)

func (p Priority) String() string {
	switch p {
	case High:   return "High"
	case Medium: return "Medium"
	default:     return "Low"
	}
}

type Task struct {
	ID        int       `json:"id"`
	Title     string    `json:"title"`
	Done      bool      `json:"done"`
	Priority  Priority  `json:"priority"`
	CreatedAt time.Time `json:"created_at"`
}