package task

import (
	"time"
)

// checks if a task should be deleted based on EOD rules
func (t *Item) IsExpired() bool {
	if !t.DeleteByEOD {
		return false
	}

	taskDate := t.TimeCreated.Format("2006-01-02")
	today := time.Now().Format("2006-01-02")
	return taskDate != today
}
