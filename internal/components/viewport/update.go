package viewport

import (
	"charm.land/bubbles/v2/key"
	tea "charm.land/bubbletea/v2"
)

// Update handles standard message-based viewport updates.
func (m Model) Update(msg tea.Msg) (Model, tea.Cmd) {
	m = m.updateAsModel(msg)
	return m, nil
}

// Author's note: this method has been broken out to make it easier to
// potentially transition Update to satisfy tea.Model.
func (m Model) updateAsModel(msg tea.Msg) Model {
	if !m.initialized {
		m.setInitialValues()
	}

	switch msg := msg.(type) {
	case tea.KeyPressMsg:
		switch {
		case key.Matches(msg, m.KeyMap.PageDown):
			m.PageDown()

		case key.Matches(msg, m.KeyMap.PageUp):
			m.PageUp()

		case key.Matches(msg, m.KeyMap.HalfPageDown):
			m.HalfPageDown()

		case key.Matches(msg, m.KeyMap.HalfPageUp):
			m.HalfPageUp()

		case key.Matches(msg, m.KeyMap.Down):
			m.ScrollDown(1)

		case key.Matches(msg, m.KeyMap.Up):
			m.ScrollUp(1)

		case key.Matches(msg, m.KeyMap.Left):
			m.ScrollLeft(m.horizontalStep)

		case key.Matches(msg, m.KeyMap.Right):
			m.ScrollRight(m.horizontalStep)
		}

	case tea.MouseWheelMsg:
		if !m.MouseWheelEnabled {
			break
		}
		switch msg.Button {
		case tea.MouseWheelDown:
			// NOTE: some terminal emulators don't send the shift event for
			// mouse actions.
			if msg.Mod.Contains(tea.ModShift) {
				m.ScrollRight(m.horizontalStep)
				break
			}
			m.ScrollDown(m.MouseWheelDelta)
		case tea.MouseWheelUp:
			// NOTE: some terminal emulators don't send the shift event for
			// mouse actions.
			if msg.Mod.Contains(tea.ModShift) {
				m.ScrollLeft(m.horizontalStep)
				break
			}
			m.ScrollUp(m.MouseWheelDelta)
		case tea.MouseWheelLeft:
			m.ScrollLeft(m.horizontalStep)
		case tea.MouseWheelRight:
			m.ScrollRight(m.horizontalStep)
		}
	}

	return m
}
