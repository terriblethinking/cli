package styles

import (
	"charm.land/lipgloss/v2"
)

var (
	TextareaBorderStyle = lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color("238"))
)
