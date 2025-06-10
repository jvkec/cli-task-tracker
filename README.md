# Boring Todo List

A boring CLI todo list written in Go.

```
 ____             _               _____        _        _ _     _   
| __ )  ___  _ __(_)_ __   __ _  |_   _|__  __| | ___  | (_)___| |_ 
|  _ \ / _ \| '__| | '_ \ / _' |   | |/ _ \/ _' |/ _ \ | | / __| __|
| |_) | (_) | |  | | | | | (_| |   | | (_) | (_| | (_) || |\__ \ |_ 
|____/|\___/|_|  |_|_| |_|\__, |   |_|\___/\__,_|\___/_|_|___/\__|
                          |___/                                      

so you can never leave your terminal #i_love_my_terminal
----------------------------------------------------------------
```

## Features

- Simple and efficient task management
- Tasks automatically delete at end of day (EOD)
- Real-time task statistics and progress tracking
- Persistent JSON storage for reliable data management
- Thread-safe operations for data integrity
- Daily completion tracking and analytics
- Flexible task prioritization (high, medium, low)

## Installation & Running

You have two options for running the todo list:

### 1. Local Build (Development)
```bash
# Clone the repository
git clone https://github.com/jvkec/cli-task-tracker.git
cd cli-task-tracker

# Build locally
make build

# Run the local build
./bin/btl --help
```

### 2. System-wide Installation
```bash
# Clone the repository
git clone https://github.com/jvkec/cli-task-tracker.git
cd cli-task-tracker

# Build and install (may require sudo)
make install

# Now you can run from anywhere
btl --help

# To uninstall later
make uninstall
```

## Usage

The todo list provides straightforward commands for managing your tasks:

```bash
# Add a temporary task (will be deleted at EOD)
btl add -t "Quick task"

# Add a permanent task (won't be deleted at EOD)
btl add -t "Important task" --keep

# Add a task with all options
btl add -t "Full task" -n "Task details" -p high --keep

# List all tasks
btl list
btl ls      # Short version

# Complete a task
btl complete 1
btl c 1     # Short version

# Delete a task (or multiple tasks)
btl delete 1
btl delete 1 2 3   # Delete multiple tasks
btl rm 1           # Short version

# View task statistics
btl stats
btl s       # Short version
```

### Command Reference

#### Adding Tasks
```bash
btl add -t "Task title" [-n "Notes"] [-p high|medium|low] [--keep]
```
Creates a new task. By default, tasks are deleted at EOD unless `--keep` is specified.

Options:
- `-t, --title`: Task title (required)
- `-n, --notes`: Additional notes (optional)
- `-p, --priority`: Task priority (optional, default: medium)
- `--keep, -k`: Keep task after EOD (optional, default: false)

#### Listing Tasks
```bash
btl list    # Show all tasks
btl ls      # Short version
```
Displays all tasks with their IDs, titles, priorities, and completion status.
Automatically cleans up expired tasks from previous days.

#### Completing Tasks
```bash
btl complete <task-id>   # Mark as done
btl c <task-id>          # Short version
```
Marks the specified task as completed.

#### Deleting Tasks
```bash
btl delete <task-id>     # Remove task
btl rm <task-id>         # Short version
```
Permanently removes the specified task.

#### Viewing Statistics
```bash
btl stats   # View completion stats
btl s       # Short version
```
Shows task completion rates and progress over time, including daily breakdowns.

## Task Lifecycle

- Tasks are temporary by default and will be deleted at the end of the day
- Use the `--keep` flag when adding a task to make it permanent
- Temporary tasks from previous days are automatically cleaned up when you list tasks
- The cleanup process helps keep your task list focused on current items

## Storage

Your tasks are automatically saved to `~/.task-tracker/tasks.json`. This file is updated in real-time as you manage your tasks, ensuring your data is always preserved between sessions.

If you need to find your tasks file, it's located at:
- Unix/Linux/macOS: `~/.task-tracker/tasks.json`
- Windows: `%USERPROFILE%\.task-tracker\tasks.json`

## Development

For those interested in working on the todo list:

```bash
# Quick run during development (builds and runs immediately)
make run

# Build without installing (creates bin/btl)
make build

# Clean build artifacts
make clean
```

## License

[MIT](LICENSE)
