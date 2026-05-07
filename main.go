package main

import (
	"fmt"
	"os"
)

func main() {
	repo := NewFileRepository("tasks.json")
	svc  := NewTaskService(repo)
	cli  := NewCLI(svc)

	if len(os.Args) < 2 {
		cli.PrintHelp()
		os.Exit(0)
	}

	if err := cli.Run(os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, "Error:", err)
		os.Exit(1)
	}
}