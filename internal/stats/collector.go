package stats

import (
	"sync/atomic"
	"time"

	"github.com/jvkec/cli-task-tracker/internal/task"
)

// handles task statistics collection
type Collector struct {
	totalTasks     atomic.Int64
	completedTasks atomic.Int64
	dailyStats     map[string]*DailyStats
}

// track statistics for a single day
type DailyStats struct {
	Date           time.Time
	Total          atomic.Int64
	Completed      atomic.Int64
	CompletionRate float64
}

// create a new statistics collector
func NewCollector() *Collector {
	return &Collector{
		dailyStats: make(map[string]*DailyStats),
	}
}

// updates statistics based on the provided tasks
func (c *Collector) UpdateStats(tasks []task.Item) {
	c.totalTasks.Store(0)
	c.completedTasks.Store(0)

	// group by date
	dailyMap := make(map[string]*DailyStats)

	for _, t := range tasks {
		date := t.TimeCreated.Format("2006-01-02")
		if _, exists := dailyMap[date]; !exists {
			dailyMap[date] = &DailyStats{
				Date: t.TimeCreated,
			}
		}

		dailyMap[date].Total.Add(1)
		c.totalTasks.Add(1)

		if t.Completed {
			dailyMap[date].Completed.Add(1)
			c.completedTasks.Add(1)
		}
	}

	// calc completion rates
	for _, stats := range dailyMap {
		total := float64(stats.Total.Load())
		if total > 0 {
			stats.CompletionRate = float64(stats.Completed.Load()) / total * 100
		}
	}

	c.dailyStats = dailyMap
}

// returns the overall task statistics
func (c *Collector) GetOverallStats() (total, completed int64, rate float64) {
	total = c.totalTasks.Load()
	completed = c.completedTasks.Load()
	if total > 0 {
		rate = float64(completed) / float64(total) * 100
	}
	return
}

// returns statistics grouped by day
func (c *Collector) GetDailyStats() map[string]*DailyStats {
	return c.dailyStats
}
