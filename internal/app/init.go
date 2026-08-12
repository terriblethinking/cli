package app

import (
	"github.com/terriblethinking/cli/internal/screens/chat"

	tea "charm.land/bubbletea/v2"
	bifrost "github.com/maximhq/bifrost/core"
)

func New(client bifrost.Bifrost) Model {
	return Model{
		chat: chat.New(client),
	}
}

func (m Model) Init() tea.Cmd {
	return tea.Batch(
		m.chat.Init(),
	)
}
