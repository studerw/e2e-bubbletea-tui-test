// Package app contains the main application logic and model.
package app

import (
	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"

	"github.com/yourusername/tui-e2e-demo/internal/components/footer"
	"github.com/yourusername/tui-e2e-demo/internal/components/header"
	"github.com/yourusername/tui-e2e-demo/internal/components/tabs"
	"github.com/yourusername/tui-e2e-demo/internal/panels/text"
	"github.com/yourusername/tui-e2e-demo/internal/panels/timer"
	"github.com/yourusername/tui-e2e-demo/internal/panels/todolist"
	"github.com/yourusername/tui-e2e-demo/internal/panels/webcall"
)

// Model represents the main application state.
type Model struct {
	header      header.Model
	tabs        tabs.Model
	footer      footer.Model
	textPanel   text.Model
	webPanel    webcall.Model
	todoPanel   todolist.Model
	timerPanel  timer.Model
	activeTab   int
	width       int
	height      int
	initialized bool
}

// New creates a new application model.
func New() Model {
	tabNames := []string{"Text Panel", "Web Call", "Todo List", "Timer"}

	return Model{
		header:     header.New("E2E MVP using BubbleTea v2"),
		tabs:       tabs.New(tabNames),
		footer:     footer.New("© 2024 Clarity Innovations™ • All Rights Reserved"),
		textPanel:  text.New(),
		webPanel:   webcall.New(),
		todoPanel:  todolist.New(),
		timerPanel: timer.New(),
		activeTab:  0,
	}
}

// Init initializes the application.
// In BubbleTea v2, Init returns (tea.Model, tea.Cmd)
func (m Model) Init() tea.Cmd {
	return nil
}

// Update handles messages and updates the model.
func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyPressMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "tab":
			m.activeTab = (m.activeTab + 1) % 4
			m.tabs.SetActive(m.activeTab)
		case "shift+tab":
			m.activeTab = (m.activeTab - 1 + 4) % 4
			m.tabs.SetActive(m.activeTab)
		case "1":
			m.activeTab = 0
			m.tabs.SetActive(m.activeTab)
		case "2":
			m.activeTab = 1
			m.tabs.SetActive(m.activeTab)
		case "3":
			m.activeTab = 2
			m.tabs.SetActive(m.activeTab)
		case "4":
			m.activeTab = 3
			m.tabs.SetActive(m.activeTab)
		}

	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
		m.initialized = true

		headerHeight := int(float64(m.height) * 0.10)
		footerHeight := int(float64(m.height) * 0.10)
		contentHeight := m.height - headerHeight - footerHeight - 4

		m.header.SetSize(m.width, headerHeight)
		m.tabs.SetSize(m.width, 3)
		m.footer.SetSize(m.width, footerHeight)

		m.textPanel.SetSize(m.width, contentHeight)
		m.webPanel.SetSize(m.width, contentHeight)
		m.todoPanel.SetSize(m.width, contentHeight)
		m.timerPanel.SetSize(m.width, contentHeight)
	}

	var cmd tea.Cmd
	switch m.activeTab {
	case 0:
		m.textPanel, cmd = m.textPanel.Update(msg)
		cmds = append(cmds, cmd)
	case 1:
		m.webPanel, cmd = m.webPanel.Update(msg)
		cmds = append(cmds, cmd)
	case 2:
		m.todoPanel, cmd = m.todoPanel.Update(msg)
		cmds = append(cmds, cmd)
	case 3:
		m.timerPanel, cmd = m.timerPanel.Update(msg)
		cmds = append(cmds, cmd)
	}

	return m, tea.Batch(cmds...)
}

// View renders the application.
func (m Model) View() tea.View {
	if !m.initialized {
		return tea.NewView("Initializing...")
	}

	var content string
	switch m.activeTab {
	case 0:
		content = m.textPanel.View()
	case 1:
		content = m.webPanel.View()
	case 2:
		content = m.todoPanel.View()
	case 3:
		content = m.timerPanel.View()
	}

	str := lipgloss.JoinVertical(
		lipgloss.Left,
		m.header.View(),
		m.tabs.View(),
		content,
		m.footer.View(),
	)

	v := tea.NewView(str)
	v.AltScreen = true
	v.MouseMode = tea.MouseModeCellMotion
	return v
}

// Run starts the application.
func Run() error {
	p := tea.NewProgram(New())
	_, err := p.Run()
	return err
}
