package main

import (
	"fmt"
	"strconv"
	"strings"
)

type CLI struct{ svc *TaskService }

func NewCLI(s *TaskService) *CLI { return &CLI{svc: s} }

func (c *CLI) Run(args []string) error {
	switch args[0] {
	case "add":
		if len(args) < 2 { return fmt.Errorf("usage: add <title> [low|medium|high]") }
		p := parsePriority(args)
		title := strings.Join(args[1:len(args)-1], " ")
		if p == Low && len(args) == 2 { title = args[1] }
		return c.svc.Add(title, p)
	case "done":
		if len(args) < 2 { return fmt.Errorf("usage: done <id>") }
		id, _ := strconv.Atoi(args[1])
		return c.svc.Complete(id)
	case "delete":
		if len(args) < 2 { return fmt.Errorf("usage: delete <id>") }
		id, _ := strconv.Atoi(args[1])
		return c.svc.Delete(id)
	case "list":
		return c.printList()
	default:
		c.PrintHelp()
	}
	return nil
}

func (c *CLI) printList() error {
	tasks, err := c.svc.List()
	if err != nil { return err }
	if len(tasks) == 0 { fmt.Println("No tasks yet!"); return nil }
	fmt.Printf("%-4s %-30s %-8s %s\n", "ID", "Title", "Priority", "Status")
	fmt.Println(strings.Repeat("-", 55))
	for _, t := range tasks {
		status := "[ ]"
		if t.Done { status = "[✓]" }
		fmt.Printf("%-4d %-30s %-8s %s\n", t.ID, t.Title, t.Priority, status)
	}
	return nil
}

func (c *CLI) PrintHelp() {
	fmt.Println("GoTask - CLI Task Manager")
	fmt.Println("Commands:")
	fmt.Println("1. add <title> [low|medium|high]  [Add a task]")
	fmt.Println("2. list                           [List all tasks]")
	fmt.Println("3. done <id>                      [Mark task as done]")
	fmt.Println("4. delete <id>                    [Delete a task]")
	fmt.Println("To run, use: go run . <command> [args]")
}

func parsePriority(args []string) Priority {
	switch strings.ToLower(args[len(args)-1]) {
	case "high":   return High
	case "medium": return Medium
	default:       return Low
	}
}