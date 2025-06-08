// No need for accessors or mutators. Direct access is idiomatic since the program is
// reading/writing a JSON file.
package task

import "time"

type Priority int

const (
	High Priority = iota
	Medium
	Low
)

// Priority to string conversion
func (p Priority) String() string {
	return [...]string{"High", "Medium", "Low"}[p]
}

type Item struct {
	ID          int       `json:"id"`
	Title       string    `json:"title"`
	Prio        Priority  `json:"priority"`
	Notes       string    `json:"notes"`
	Completed   bool      `json:"completed"`
	DeleteByEOD bool      `json:"delete_by_eod"`
	TimeCreated time.Time `json:"time_created"`
}
