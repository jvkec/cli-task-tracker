package app

import (
	"github.com/jvkec/cli-task-tracker/internal/cli"
	ucli "github.com/urfave/cli/v2"
)

func New() *ucli.App {
	app := &ucli.App{
		Name:     "task-tracker",
		Usage:    "A boring terminal todo list",
		Commands: cli.GetCommands(),
	}
	return app
}

func addTask(c *ucli.Context) error {
	// TODO: Implement add task
	return nil
}

func listTasks(c *ucli.Context) error {
	// TODO: Implement list tasks
	return nil
}

func completeTask(c *ucli.Context) error {
	// TODO: Implement complete task
	return nil
}

func deleteTask(c *ucli.Context) error {
	// TODO: Implement delete task
	return nil
}
