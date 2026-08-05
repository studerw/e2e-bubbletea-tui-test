// Package footer provides the footer component for the TUI.
package footer

import (
	"charm.land/lipgloss/v2"
)

var (
	footerStyle = lipgloss.NewStyle().
		Foreground(lipgloss.Color("#ABABAB")).
		Background(lipgloss.Color("#333333")).
		Padding(0, 2).
		Align(lipgloss.Center)
)

// Model represents the footer component state.
type Model struct {
	text   string
	width  int
	height int
}

// New creates a new footer model.
func New(text string) Model {
	return Model{
		text: text,
	}
}

// SetSize sets the footer dimensions.
func (m *Model) SetSize(width, height int) {
	m.width = width
	m.height = height
}

// View renders the footer.
func (m Model) View() string {
	helpText := " • Press 1-4 to switch tabs • Tab/Shift+Tab to navigate • q to quit"
	return footerStyle.
		Width(m.width).
		Height(m.height).
		Render(m.text + helpText)
}
