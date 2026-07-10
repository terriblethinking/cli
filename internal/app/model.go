package app

import (
	tea "charm.land/bubbletea/v2"
)

type Screen int

const (
	ChatScreen Screen = iota
)

type Model struct {
	current Screen

	width  int
	height int

	chat tea.Model
}
