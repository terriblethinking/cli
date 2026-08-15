package chat

import (
	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
	"strings"
)

func (m Model) messagesView() string {
	var messages strings.Builder

	for _, message := range m.turns {
		if message.Role == "person" {
			messages.WriteString(userMessageStyle.Width(m.width).Render(message.Content))
		}

		if message.Role == "model-thinking" {
			messages.WriteString(modelThinkingStyle.Width(m.width).Render(message.Content))
		}

		if message.Role == "model-text" {
			messages.WriteString(modelTextStyle.Width(m.width).Render(message.Content))
		}
	}

	if m.incomingThinking.String() != "" {
		messages.WriteString(modelThinkingStyle.Width(m.width).Render(m.incomingThinking.String()))
	}

	if m.incomingText.String() != "" {
		messages.WriteString(modelTextStyle.Width(m.width).Render(m.incomingText.String()))
	}

	return messages.String()
}

func (m Model) View() tea.View {

	// ui := strings.Builder{}

	const (
		footer = "\n(ctrl+c to quit)\n"
	)

	var c *tea.Cursor
	if !m.textarea.VirtualCursor() {
		c = m.textarea.Cursor()

		if c != nil {
			// Set the y offset of the cursor based on the position of the textarea
			// in the application.
			offset := lipgloss.Height(m.messagesView())
			c.Y += offset
		}
	}

	f := strings.Join([]string{
		m.messagesView(),
		m.textarea.View(),
		footer,
	}, "\n")

	v := tea.NewView(f)
	v.Cursor = c
	return v
}
