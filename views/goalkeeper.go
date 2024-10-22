package views

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var (
	goalkeeperPosition = lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("#D8315B")).Render("⬤")
)

type GoalkeeperModel struct {
	y          int
	rangeInit  int
	rangeFinal int
}

func NewGoalkeeperModel(rangeInit, rangeFinal int) GoalkeeperModel {
	return GoalkeeperModel{
		rangeInit:  rangeInit,
		rangeFinal: rangeFinal,
		y:          (rangeInit + rangeFinal) / 2,
	}
}

func (m GoalkeeperModel) Init() tea.Cmd {
	return nil
}

func (m GoalkeeperModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "up":
			if m.y > m.rangeInit {
				m.y--
			}
		case "down":
			if m.y < m.rangeFinal {
				m.y++
			}
		}
	}
	return m, nil
}

func (m GoalkeeperModel) View() string {
	return goalkeeperPosition
}
