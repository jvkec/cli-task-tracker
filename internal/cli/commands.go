package cli

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/jvkec/cli-task-tracker/internal/stats"
	"github.com/jvkec/cli-task-tracker/internal/storage"
	"github.com/jvkec/cli-task-tracker/internal/task"
	ucli "github.com/urfave/cli/v2"
)

// task storage interface
var store *storage.Store

// returns all available cli cmds
func GetCommands(s *storage.Store) []*ucli.Command {
	store = s
	return []*ucli.Command{
		AddCommand(),
		ListCommand(),
		DeleteCommand(),
		CompleteCommand(),
		StatsCommand(),
	}
}

// converts a string priority to task.Priority
func parsePriority(p string) task.Priority {
	switch strings.ToLower(p) {
	case "high":
		return task.High
	case "low":
		return task.Low
	default:
		return task.Medium
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
			&ucli.BoolFlag{
				Name:    "keep",
				Aliases: []string{"k"},
				Usage:   "Keep task after EOD (by default tasks are deleted at EOD)",
				Value:   false,
			},
		},
		Action: func(ctx *ucli.Context) error {
			t := task.Item{
				Title:       ctx.String("title"),
				Notes:       ctx.String("notes"),
				Prio:        parsePriority(ctx.String("priority")),
				TimeCreated: time.Now(),
				DeleteByEOD: !ctx.Bool("keep"), // Default to true unless --keep is specified
			}

			newTask, err := store.Add(t)
			if err != nil {
				return fmt.Errorf("failed to add task: %w", err)
			}

			fmt.Printf("Added task %d: %s (Priority: %s)\n", newTask.ID, newTask.Title, newTask.Prio)
			if newTask.DeleteByEOD {
				fmt.Println("Note: This task will be deleted at end of day. Use --keep flag to preserve it.")
			}
			return nil
		},
	}
}

// Show all tasks
func ListCommand() *ucli.Command {
	return &ucli.Command{
		Name:    "list",
		Aliases: []string{"l", "ls"},
		Usage:   "List all tasks",
		Action: func(ctx *ucli.Context) error {
			tasks := store.List()
			if len(tasks) == 0 {
				fmt.Println("No tasks found")
				return nil
			}

			fmt.Println("\nTasks:")
			fmt.Println("-------")
			for _, t := range tasks {
				status := "[ ]"
				if t.Completed {
					status = "[✓]"
				}
				fmt.Printf("%s #%d: %s (Priority: %s)\n", status, t.ID, t.Title, t.Prio)
				if t.Notes != "" {
					fmt.Printf("    Notes: %s\n", t.Notes)
				}
			}
			fmt.Println()
			return nil
		},
	}
}

// Remove a task
func DeleteCommand() *ucli.Command {
	return &ucli.Command{
		Name:      "delete",
		Aliases:   []string{"d", "rm"},
		Usage:     "Delete one or more tasks by ID",
		ArgsUsage: "TASK_ID [TASK_ID...]",
		Action: func(ctx *ucli.Context) error {
			if !ctx.Args().Present() {
				return fmt.Errorf("at least one task ID is required")
			}

			var failedDeletes []string
			for _, arg := range ctx.Args().Slice() {
				id, err := strconv.Atoi(arg)
				if err != nil {
					failedDeletes = append(failedDeletes, fmt.Sprintf("%s (invalid ID)", arg))
					continue
				}

				if err := store.Delete(id); err != nil {
					failedDeletes = append(failedDeletes, fmt.Sprintf("%d (%v)", id, err))
					continue
				}

				fmt.Printf("Deleted task %d\n", id)
			}

			if len(failedDeletes) > 0 {
				return fmt.Errorf("failed to delete some tasks: %s", strings.Join(failedDeletes, ", "))
			}

			return nil
		},
	}
}

// Mark a task as completed
func CompleteCommand() *ucli.Command {
	return &ucli.Command{
		Name:      "complete",
		Aliases:   []string{"c", "done"},
		Usage:     "Mark a task as completed",
		ArgsUsage: "TASK_ID",
		Action: func(ctx *ucli.Context) error {
			if !ctx.Args().Present() {
				return fmt.Errorf("task ID is required")
			}

			id, err := strconv.Atoi(ctx.Args().First())
			if err != nil {
				return fmt.Errorf("invalid task ID: %s", ctx.Args().First())
			}

			task, exists := store.Get(id)
			if !exists {
				return fmt.Errorf("task %d not found", id)
			}

			task.Completed = true
			if err := store.Update(task); err != nil {
				return fmt.Errorf("failed to update task: %w", err)
			}

			fmt.Printf("Marked task %d as completed\n", id)
			return nil
		},
	}
}

// show task statistics
func StatsCommand() *ucli.Command {
	return &ucli.Command{
		Name:    "stats",
		Aliases: []string{"s"},
		Usage:   "Show task statistics",
		Action: func(ctx *ucli.Context) error {
			tasks := store.List()
			collector := stats.NewCollector()
			collector.UpdateStats(tasks)

			total, completed, _ := collector.GetOverallStats()
			fmt.Printf("\nOverall Statistics:\n")
			fmt.Printf("Total Tasks: %d\n", total)
			fmt.Printf("Completed:   %d\n", completed)
			fmt.Printf("Progress:    %s\n\n", stats.RenderProgressBar(completed, total))

			dailyStats := collector.GetDailyStats()
			if len(dailyStats) > 0 {
				fmt.Println("Daily Statistics:")
				for date, ds := range dailyStats {
					fmt.Printf("\n%s:\n", date)
					fmt.Printf("  Tasks:      %d (Completed: %d)\n", ds.Total.Load(), ds.Completed.Load())
					fmt.Printf("  Progress:   %s\n", stats.RenderProgressBar(ds.Completed.Load(), ds.Total.Load()))
				}
			}

			return nil
		},
	}
}
