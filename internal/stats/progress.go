package stats

import (
	"fmt"
	"strings"
)

const (
	progressBarWidth = 30
	progressChar     = "█"
	emptyChar        = "░"
)

// visual progress bar creation
func RenderProgressBar(completed, total int64) string {
	if total == 0 {
		return fmt.Sprintf("[%s] 0%%", strings.Repeat(emptyChar, progressBarWidth))
	}

	percentage := float64(completed) / float64(total)
	filledWidth := int(float64(progressBarWidth) * percentage)

	if filledWidth > progressBarWidth {
		filledWidth = progressBarWidth
	}

	emptyWidth := progressBarWidth - filledWidth

	bar := fmt.Sprintf("[%s%s] %.1f%%",
		strings.Repeat(progressChar, filledWidth),
		strings.Repeat(emptyChar, emptyWidth),
		percentage*100,
	)

	return bar
}
