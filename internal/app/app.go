package app

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/jvkec/cli-task-tracker/internal/cli"
	"github.com/jvkec/cli-task-tracker/internal/storage"
	ucli "github.com/urfave/cli/v2"
)

const banner = `
_                _               _            _         _ _     _   
| |__   ___  _ __(_)_ __   __ _  | |_ ___   __| | ___   | (_)___| |_ 
| '_ \ / _ \| '__| | '_ \ / _` | | __/ _ \ / _` |/ _ \  | | / __| __|
| |_) | (_) | |  | | | | | (_| | | || (_) | (_| | (_) | | | \__ \ |_ 
|_.__/ \___/|_|  |_|_| |_|\__, |  \__\___/ \__,_|\___/  |_|_|___/\__|
						  |___/                                      

so you can never leave your terminal #i_love_my_terminal
----------------------------------------------------------------
`

const helpText = `
USAGE:
   btl [command] [options]

COMMANDS:
   add, a        Add a new task (--title, --notes, --priority, --keep)
   list, ls      List all tasks
   complete, c   Mark a task as complete
   delete, rm    Delete one or more tasks by ID
   stats, s      Show task statistics
   help, h       Show this help

FLAGS:
   --help, -h     Show help
   --version, -v  Print version

EXAMPLES:
   btl add -t "Important task" -n "Details here" -p high
   btl add -t "Temporary task" (will be deleted at EOD)
   btl add -t "Persistent task" --keep (won't be deleted at EOD)
   btl list
   btl complete 1
   btl delete 1 2 3
   btl stats
`

// New creates and returns a new CLI application instance
func New() *ucli.App {
	// Setup storage
	homeDir, err := os.UserHomeDir()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Warning: Could not get home directory: %v\n", err)
		fmt.Fprintf(os.Stderr, "Using current directory for storage\n")
		homeDir = "."
	}
	storageDir := filepath.Join(homeDir, ".task-tracker")
	storePath := filepath.Join(storageDir, "tasks.json")
	store, err := storage.NewStore(storePath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to initialize storage at %s: %v\n", storePath, err)
		os.Exit(1)
	}

	app := &ucli.App{
		Name:        "btl",
		Usage:       "A boring todo list",
		Description: strings.TrimSpace(banner) + "\n\n" + strings.TrimSpace(helpText),
		Commands:    cli.GetCommands(store),
		Authors: []*ucli.Author{
			{
				Name: "Boring Todo List",
			},
		},
		Version: "1.0.0",
		Action: func(c *ucli.Context) error {
			// show banner and help when no command is provided
			fmt.Print(banner)
			fmt.Print(helpText)
			return nil
		},
	}

	return app
}
