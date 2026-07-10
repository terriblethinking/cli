package app

import tea "charm.land/bubbletea/v2"

func (m Model) View() tea.View {
	switch m.current {
	case ChatScreen:
		return m.chat.View()
	}

	return tea.NewView("")
}
