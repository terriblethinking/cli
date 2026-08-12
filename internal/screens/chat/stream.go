package chat

import (
	tea "charm.land/bubbletea/v2"
)

type StreamDoneMsg struct{}

func recieveChunk(ch <-chan any) tea.Cmd {
	return func() tea.Msg {
		chunk, ok := <-ch
		if !ok {
			return StreamDoneMsg{}
		}
		return chunk
	}
}
