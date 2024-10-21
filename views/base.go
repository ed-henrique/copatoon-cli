package views

import tea "github.com/charmbracelet/bubbletea"

type baseModel struct {
	windowWidth  int
	windowHeight int
	currentModel tea.Model
}

func NewBaseModel() baseModel {
	initialModel := NewStartModel()

	return baseModel{
		currentModel: &initialModel,
	}
}

func (m baseModel) Init() tea.Cmd {
	return m.currentModel.Init()
}

func (m baseModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	return m.currentModel.Update(msg)
}

func (m baseModel) View() string {
	return m.currentModel.View()
}

func (m baseModel) SwitchView(model tea.Model) (tea.Model, tea.Cmd) {
	clearTerminal()
	m.currentModel = model
	return m.currentModel, m.currentModel.Init()
}
