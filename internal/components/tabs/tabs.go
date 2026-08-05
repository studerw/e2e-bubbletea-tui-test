// Package tabs provides the tab navigation component.
package tabs

import (
	"charm.land/lipgloss/v2"
)

var (
	activeTabStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("#FFFFFF")).
			Background(lipgloss.Color("#7D56F4")).
			Padding(0, 2)

	inactiveTabStyle = lipgloss.NewStyle().
				Foreground(lipgloss.Color("#888888")).
				Background(lipgloss.Color("#333333")).
				Padding(0, 2)

	tabGap = lipgloss.NewStyle().Width(1)
)

// Model represents the tabs component state.
type Model struct {
	tabs      []string
	activeTab int
	width     int
	height    int
}

// New creates a new tabs model.
func New(tabs []string) Model {
	return Model{
		tabs:      tabs,
		activeTab: 0,
	}
}

// SetActive sets the active tab index.
func (m *Model) SetActive(index int) {
	if index >= 0 && index < len(m.tabs) {
		m.activeTab = index
	}
}

// Active returns the currently active tab index.
func (m Model) Active() int {
	return m.activeTab
}

// SetSize sets the tabs dimensions.
func (m *Model) SetSize(width, height int) {
	m.width = width
	m.height = height
}

// View renders the tabs.
func (m Model) View() string {
	var renderedTabs []string

	for i, tab := range m.tabs {
		shortcut := string(rune('1' + i))
		tabLabel := "[" + shortcut + "] " + tab

		if i == m.activeTab {
			renderedTabs = append(renderedTabs, activeTabStyle.Render(tabLabel))
		} else {
			renderedTabs = append(renderedTabs, inactiveTabStyle.Render(tabLabel))
		}
	}

	// Join tabs horizontally with a small gap between them
	row := lipgloss.JoinHorizontal(lipgloss.Top, renderedTabs...)

	// Center the tabs row within the available width
	centered := lipgloss.NewStyle().
		Width(m.width).
		Align(lipgloss.Center).
		Background(lipgloss.Color("#1A1A1A")).
		Padding(1, 0).
		Render(row)

	return centered
}
