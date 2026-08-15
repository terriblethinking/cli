package chat

import (
	"fmt"

	"charm.land/bubbles/v2/textarea"
	tea "charm.land/bubbletea/v2"
	"github.com/terriblethinking/engine/agent"
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

	case agent.AsyncAgentReasongingChunk:

		if m.state != "receiving" {
			m.state = "receiving"
		}

		if m.incomingText.String() != "" {

			m.turns = append(m.turns, Turn{
				Role:    "model-text",
				Content: m.incomingText.String(),
			})

			m.incomingText.Reset()
		}

		m.incomingThinking.WriteString(msg.Content)

		return m, recieveChunk(m.streamCh)

	case agent.AsyncAgentTextChunk:

		if m.state != "receiving" {
			m.state = "receiving"
		}

		if m.incomingThinking.String() != "" {

			m.turns = append(m.turns, Turn{
				Role:    "model-thinking",
				Content: m.incomingThinking.String(),
			})

			m.incomingThinking.Reset()
		}

		m.incomingText.WriteString(msg.Content)

		return m, recieveChunk(m.streamCh)

	case StreamDoneMsg:

		if m.incomingThinking.String() != "" {

			m.turns = append(m.turns, Turn{
				Role:    "model-thinking",
				Content: m.incomingThinking.String(),
			})

			m.incomingThinking.Reset()
		}

		if m.incomingText.String() != "" {

			m.turns = append(m.turns, Turn{
				Role:    "model-text",
				Content: m.incomingText.String(),
			})

			m.incomingText.Reset()
		}

		m.state = "idle"

	case tea.KeyPressMsg:
		switch msg.String() {
		case "esc":
			if m.textarea.Focused() {
				m.textarea.Blur()
			}
		case "enter":
			if m.state != "receiving" && m.state != "loading" {
				if m.textarea.Focused() {
					m.state = "loading"

					// Get the current user prompt.

					prompt := m.textarea.Value()

					// Add the user prompt to the turn slice

					m.turns = append(m.turns, Turn{
						Role:    "person",
						Content: prompt,
					})

					// Reset the value of the input to "".

					m.textarea.SetValue("")

					// Get the agent readable messages we feed the agent.

					messages := m.GetAgentReadableMessages()

					fmt.Print(len(messages))

					// Set them inside the agent

					m.agent.Messages = messages

					// Start the streaming stream and get the channel.

					streamCh := m.agent.RunAsync(prompt)

					// Save the channel inside the model so that we
					// can access it later on.

					m.streamCh = streamCh

					// Return recieveChunk that will continue to propagate
					// everything that it recieves from the stream channel.

					return m, recieveChunk(streamCh)
				}
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
