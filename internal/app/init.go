package app

import (
	tea "charm.land/bubbletea/v2"
	"cli/internal/screens/chat"
)

func New() Model {
	return Model{
		chat: chat.New(),
	}
}

func (m Model) Init() tea.Cmd {
	return tea.Batch(
		m.chat.Init(),
	)
}
