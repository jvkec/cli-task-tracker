package cli

import (
	"fmt"

	ucli "github.com/urfave/cli/v2"
	// "github.com/jvkec/cli-task-tracker/internal/task"
)

// Return all commands
func GetCommands() []*ucli.Command {
	return []*ucli.Command{
		AddCommand(),
		ListCommand(),
		DeleteCommand(),
		CompleteCommand(),
	}
}

// Create a new task
func AddCommand() *ucli.Command {
	return &ucli.Command{
		Name:    "add",
		Aliases: []string{"a"},
		Usage:   "Add a new task",
		Flags: []ucli.Flag{
			&ucli.StringFlag{
				Name:     "title",
				Aliases:  []string{"t"},
				Usage:    "Task title",
				Required: true,
			},
			&ucli.StringFlag{
				Name:    "notes",
				Aliases: []string{"n"},
				Usage:   "Task notes",
			},
			&ucli.StringFlag{
				Name:    "priority",
				Aliases: []string{"p"},
				Usage:   "Task priority (high, medium, low)",
				Value:   "medium",
			},
		},
		Action: func(ctx *ucli.Context) error {
			// TODO: Implement add logic
			fmt.Printf("Adding task: %s\n", ctx.String("title"))
			return nil
		},
	}
}

// Show all tasks
func ListCommand() *ucli.Command {
	return &ucli.Command{
		Name:		 "list",
		Aliases: []string{"l", "ls"},
		Usage:   "List all tasks",
		Action: func(ctx *ucli.Context) error {
			// TODO: Implement list logic
			fmt.Println("Listing tasks...")
			return nil
		},
	}
}

// Remove a task
func DeleteCommand() *ucli.Command {
	return &ucli.Command{
		Name:    "delete",
		Aliases: []string{"d", "rm"},
		Usage:   "Delete a task by ID",
		Action: func(ctx *ucli.Context) error {
			// TODO: Implement delete logic
			fmt.Printf("Deleting task with ID: %s\n", ctx.Args().First())
			return nil
		},
	}
}

// Mark a task as completed
func CompleteCommand() *ucli.Command {
	return &ucli.Command{
		Name:    "complete",
		Aliases: []string{"c", "done"},
		Usage:   "Mark a task as completed",
		Action: func(ctx *ucli.Context) error {
			// TODO: Implement complete logic
			fmt.Printf("Completing task with ID: %s\n", ctx.Args().First())
			return nil
		},
	}
}
