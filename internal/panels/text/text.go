// Package text provides the text display panel with Lorem Ipsum content.
package text

import (
	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
)

const loremIpsum = `Lorem ipsum dolor sit amet, consectetur adipiscing elit. Sed do eiusmod 
tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, 
quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat.`

var (
	panelStyle = lipgloss.NewStyle().
			Padding(1, 2).
			Border(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color("#7D56F4"))

	titleStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("#7D56F4")).
			MarginBottom(1)

	textStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#CCCCCC"))
)

// Model represents the text panel state.
type Model struct {
	width  int
	height int
}

// New creates a new text panel model.
func New() Model {
	return Model{}
}

// SetSize sets the panel dimensions.
func (m *Model) SetSize(width, height int) {
	m.width = width - 4
	m.height = height - 4
}

// Update handles messages for the text panel.
func (m Model) Update(msg tea.Msg) (Model, tea.Cmd) {
	return m, nil
}

// View renders the text panel.
func (m Model) View() string {
	title := titleStyle.Render("📄 Lorem Ipsum")
	content := textStyle.Render(loremIpsum)

	return panelStyle.
		Width(m.width).
		Height(m.height).
		Render(lipgloss.JoinVertical(lipgloss.Left, title, content))
}
