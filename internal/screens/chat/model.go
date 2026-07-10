package chat

import (
	"charm.land/bubbles/v2/textarea"
	tea "charm.land/bubbletea/v2"
	"cli/internal/styles"
)

type errMsg error

type Model struct {
	textarea textarea.Model

	height int
	width  int

	err error
}

func New() Model {
	ta := textarea.New()

	ta.ShowLineNumbers = false
	ta.DynamicHeight = true
	ta.MinHeight = 1
	ta.MaxHeight = 4
	ta.SetStyles().Base(styles.TextareaBorderStyle)
	ta.Placeholder = "today i want to..."
	ta.SetVirtualCursor(false)
	ta.SetStyles(textarea.DefaultStyles(true))
	ta.Focus()

	return Model{
		textarea: ta,
		err:      nil,
	}
}

func (m Model) Init() tea.Cmd {
	return tea.Batch(
		textarea.Blink,
		tea.RequestBackgroundColor,
	)
}
