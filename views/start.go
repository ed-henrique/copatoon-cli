package views

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type startModel struct {
	choices []string
	cursor  uint8
}

func NewStartModel() startModel {
	return startModel{
		choices: []string{
			"Play",
			"Exit",
		},
	}
}

func (m startModel) Init() tea.Cmd {
	return nil
}

func (m startModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			} else {
				m.cursor = uint8(len(m.choices)) - 1
			}
		case "down", "j":
			if m.cursor < uint8(len(m.choices))-1 {
				m.cursor++
			} else {
				m.cursor = 0
			}
		case "enter", " ":
			switch m.choices[m.cursor] {
			case "Play":
				gm := NewGameModel()
				return NewBaseModel().SwitchView(&gm)
			case "Exit":
				clearTerminal()
				return m, tea.Quit
			}
		}
	}

	return m, nil
}

func (m startModel) View() string {
	s := "Select to play:\n\n"

	for i, choice := range m.choices {
		cursor := " "
		if m.cursor == uint8(i) {
			cursor = ">"

			styled := lipgloss.NewStyle().
				Foreground(lipgloss.Color("#FF0000")).
				Render(fmt.Sprintf("%s %s", cursor, choice))

			s += styled + "\n"
		} else {
			s += fmt.Sprintf("%s %s\n", cursor, choice)
		}
	}

	s += "\nPress q to quit.\n"

	return s
}
