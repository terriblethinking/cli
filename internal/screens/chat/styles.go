package chat

import (
	"charm.land/lipgloss/v2"
)

var (
	lightDark = lipgloss.LightDark(true)

	userMessageStyle = lipgloss.NewStyle().
				Align(lipgloss.Left).
				Foreground(lipgloss.Color("#FAFAFA")).
				Border(lipgloss.RoundedBorder()).
				BorderForeground(lipgloss.Color("#8b7259")).
				Margin(1, 3, 0, 0).
				Padding(0, 1)

	modelThinkingStyle = lipgloss.NewStyle().
				Align(lipgloss.Left).
				Foreground(lipgloss.Color("#717171")).
				Italic(true).
				Margin(1, 3, 0, 0).
				Padding(1, 1)

	modelTextStyle = lipgloss.NewStyle().
			Align(lipgloss.Left).
			Foreground(lipgloss.Color("#FAFAFA")).
			Margin(1, 3, 0, 0).
			Padding(1, 1)
)
