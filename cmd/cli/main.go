package main

import (
	tea "charm.land/bubbletea/v2"
	"cli/internal/app"
	"log"
)

func main() {
	p := tea.NewProgram(app.New())

	if _, err := p.Run(); err != nil {
		log.Fatal(err)
	}
}
