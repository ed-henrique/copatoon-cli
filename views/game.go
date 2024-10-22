package views

import (
	"copatoon/components"
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
	x                    int
	y                    int
	windowWidth          int
	windowHeight         int
	hasInsufficientSpace bool
}

func NewGameModel() gameModel {
	return gameModel{
		x:                    0,
		y:                    0,
		windowWidth:          80,
		windowHeight:         21,
		hasInsufficientSpace: false,
	}
}

func (m gameModel) Init() tea.Cmd {
	return tea.WindowSize()
}

func (m gameModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		if msg.Width < m.windowWidth || msg.Height < m.windowHeight {
			m.hasInsufficientSpace = true
			clearTerminal()
		} else {
			m.hasInsufficientSpace = false
		}
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
	var sb strings.Builder

	if m.hasInsufficientSpace {
		return "There's not enough space to play the game."
	}

	goal := components.Goal(m.windowWidth, m.windowHeight)

	for i := range m.windowHeight {
		for j := range m.windowWidth {
			if i == m.y && j == m.x {
				sb.WriteString(playerPosition)
			} else {
				sb.WriteString(string(goal[i][j]))
			}
		}

		if i != m.windowHeight-1 {
			sb.WriteString("\n")
		}
	}

	return lipgloss.NewStyle().Border(lipgloss.DoubleBorder()).Padding(0, 1).Render(sb.String())
}
