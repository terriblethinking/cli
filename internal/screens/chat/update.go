package chat

import (
	"charm.land/bubbles/v2/textarea"
	tea "charm.land/bubbletea/v2"
)

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd
	var cmd tea.Cmd

	switch msg := msg.(type) {

	case tea.WindowSizeMsg:
		m.height = msg.Height
		m.width = msg.Width

		m.textarea.SetWidth(msg.Width)

	case tea.BackgroundColorMsg:
		// Update styling now that we know the background color.
		m.textarea.SetStyles(textarea.DefaultStyles(msg.IsDark()))

	case tea.KeyPressMsg:
		switch msg.String() {
		case "esc":
			if m.textarea.Focused() {
				m.textarea.Blur()
			}
		case "ctrl+c":
			return m, tea.Quit
		default:
			if !m.textarea.Focused() {
				cmd = m.textarea.Focus()
				cmds = append(cmds, cmd)
			}
		}

		// We handle errors just like any other message
	case errMsg:
		m.err = msg
		return m, nil
	}

	m.textarea, cmd = m.textarea.Update(msg)
	cmds = append(cmds, cmd)
	return m, tea.Batch(cmds...)
}
