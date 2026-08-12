package chat

import (
	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
	"fmt"
	"strings"
)

func (m Model) headerView() string {
	return "Tell me a story.\n"
}

func (m Model) messagesView() string {
	var messages strings.Builder

	for _, message := range m.turns {
		messages.WriteString(fmt.Sprintf("\n%s\n\n", message.Content))
	}

	if m.incomingThinking.String() != "" {
		messages.WriteString(m.incomingThinking.String())
	}

	if m.incomingText.String() != "" {
		messages.WriteString(m.incomingText.String())
	}

	return messages.String()
}

func (m Model) View() tea.View {
	const (
		footer = "\n(ctrl+c to quit)\n"
	)

	var c *tea.Cursor
	if !m.textarea.VirtualCursor() {
		c = m.textarea.Cursor()

		if c != nil {
			// Set the y offset of the cursor based on the position of the textarea
			// in the application.
			offset := lipgloss.Height(m.headerView()) + lipgloss.Height(m.messagesView())
			c.Y += offset
		}
	}

	f := strings.Join([]string{
		m.headerView(),
		m.messagesView(),
		m.textarea.View(),
		footer,
	}, "\n")

	v := tea.NewView(f)
	v.Cursor = c
	return v
}
