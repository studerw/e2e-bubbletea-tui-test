// Package header provides the header component for the TUI.
package header

import (
	"charm.land/lipgloss/v2"
)

var (
	headerStyle = lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("#FFFFFF")).
		Background(lipgloss.Color("#7D56F4")).
		Padding(1, 2).
		Align(lipgloss.Center)
)

// Model represents the header component state.
type Model struct {
	title  string
	width  int
	height int
}

// New creates a new header model.
func New(title string) Model {
	return Model{
		title: title,
	}
}

// SetSize sets the header dimensions.
func (m *Model) SetSize(width, height int) {
	m.width = width
	m.height = height
}

// View renders the header.
func (m Model) View() string {
	return headerStyle.
		Width(m.width).
		Height(m.height).
		Render(m.title)
}
