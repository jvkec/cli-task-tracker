package storage

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"sync"

	"github.com/jvkec/cli-task-tracker/internal/task"
)

type Store struct {
	tasks    map[int]task.Item
	nextID   int
	filename string
	mu       sync.RWMutex
}

func NewStore(filename string) (*Store, error) {
	s := &Store{
		tasks:    make(map[int]task.Item),
		nextID:   1,
		filename: filename,
	}

	// create storage directory if it dont exist
	dir := filepath.Dir(filename)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return nil, fmt.Errorf("failed to create storage directory: %w", err)
	}

	// load existing tasks if file exists
	if err := s.load(); err != nil && !os.IsNotExist(err) {
		return nil, fmt.Errorf("failed to load tasks: %w", err)
	}

	return s, nil
}

func (s *Store) load() error {
	data, err := os.ReadFile(s.filename)
	if err != nil {
		return err
	}

	var tasks []task.Item
	if err := json.Unmarshal(data, &tasks); err != nil {
		return fmt.Errorf("failed to unmarshal tasks: %w", err)
	}

	s.mu.Lock()
	defer s.mu.Unlock()

	for _, t := range tasks {
		s.tasks[t.ID] = t
		if t.ID >= s.nextID {
			s.nextID = t.ID + 1
		}
	}

	return nil
}

func (s *Store) save() error {
	s.mu.RLock()
	tasks := make([]task.Item, 0, len(s.tasks))
	for _, t := range s.tasks {
		tasks = append(tasks, t)
	}
	s.mu.RUnlock()

	// marshal n save without holding lock
	data, err := json.MarshalIndent(tasks, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal tasks: %w", err)
	}

	if err := os.WriteFile(s.filename, data, 0644); err != nil {
		return fmt.Errorf("failed to write tasks file: %w", err)
	}

	return nil
}

func (s *Store) Add(t task.Item) (task.Item, error) {
	s.mu.Lock()
	t.ID = s.nextID
	s.nextID++
	s.tasks[t.ID] = t
	s.mu.Unlock()

	if err := s.save(); err != nil {
		return task.Item{}, err
	}

	return t, nil
}

func (s *Store) List() []task.Item {
	// cleanup expired tasks
	s.Cleanup()

	// get the list of remaining tasks (kept with flag by user)
	s.mu.RLock()
	tasks := make([]task.Item, 0, len(s.tasks))
	for _, t := range s.tasks {
		tasks = append(tasks, t)
	}
	s.mu.RUnlock()

	return tasks
}

func (s *Store) Get(id int) (task.Item, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	t, ok := s.tasks[id]
	return t, ok
}

func (s *Store) Update(t task.Item) error {
	s.mu.Lock()
	if _, exists := s.tasks[t.ID]; !exists {
		s.mu.Unlock()
		return fmt.Errorf("task with ID %d not found", t.ID)
	}
	s.tasks[t.ID] = t
	s.mu.Unlock()

	return s.save()
}

func (s *Store) Delete(id int) error {
	s.mu.Lock()
	if _, exists := s.tasks[id]; !exists {
		s.mu.Unlock()
		return fmt.Errorf("task with ID %d not found", id)
	}
	delete(s.tasks, id)
	s.mu.Unlock()

	return s.save()
}

func (s *Store) Cleanup() {
	// get expired tasks under read lock
	s.mu.RLock()
	expiredIDs := make([]int, 0)
	for id, t := range s.tasks {
		if t.IsExpired() {
			expiredIDs = append(expiredIDs, id)
		}
	}
	s.mu.RUnlock()

	// del expired tasks if any found
	if len(expiredIDs) > 0 {
		s.mu.Lock()
		for _, id := range expiredIDs {
			delete(s.tasks, id)
		}
		s.mu.Unlock()

		s.save()
	}
}
