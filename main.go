package main

import (
	"fmt"
	"os"

	"copatoon/views"
	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	p := tea.NewProgram(views.NewBaseModel())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Error when starting the game: %s\n", err.Error())
		os.Exit(1)
	}
}
