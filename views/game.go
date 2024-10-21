package views

import (
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type state int

const (
	initializing state = iota
	ready
)

const (
	accel = 2
)

var (
	playerPosition = lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("#9448BC")).Render("⬤")
)

type gameModel struct {
	x            int
	y            int
	windowWidth  int
	windowHeight int
}

func NewGameModel() gameModel {
	return gameModel{
		x: 0,
		y: 0,
	}
}

func (m gameModel) Init() tea.Cmd {
	return tea.WindowSize()
}

func (m gameModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.windowWidth, m.windowHeight = min(80, msg.Width-4), min(20, msg.Height-2)
		m.x = min(m.x, m.windowWidth-1)
		m.y = min(m.y, m.windowHeight-1)
		clearTerminal()
	case tea.KeyMsg:
		switch msg.String() {
		case "up":
			if m.y > 0 {
				m.y--
			}
		case "down":
			if m.y < m.windowHeight-1 {
				m.y++
			}
		case "right":
			if m.x < m.windowWidth-1 {
				m.x += accel
				m.x = min(m.x, m.windowWidth-1)
			}
		case "left":
			if m.x > 0 {
				m.x -= accel
				m.x = max(m.x, 0)
			}
		case "ctrl+c", "q":
			clearTerminal()
			return m, tea.Quit
		}
	}

	return m, nil
}

func (m gameModel) View() string {
	var (
		err error
		sb  strings.Builder
	)

	for i := range m.windowHeight {
		for j := range m.windowWidth {

			if i == m.y && j == m.x {
				_, err = sb.WriteString(playerPosition)
				if err != nil {
					panic(err)
				}
			} else {
				_, err = sb.WriteString(" ")
				if err != nil {
					panic(err)
				}
			}
		}

		if i != m.windowHeight-1 {
			sb.WriteString("\n")
		}
	}

	return lipgloss.NewStyle().Border(lipgloss.DoubleBorder()).Padding(0, 1).Render(sb.String())
}
