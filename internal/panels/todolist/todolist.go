// Package todolist provides an interactive todo list panel.
package todolist

import (
	tea "github.com/charmbracelet/bubbletea/v2"
	"github.com/charmbracelet/lipgloss/v2"
)

var (
	panelStyle = lipgloss.NewStyle().
			Padding(1, 2).
			Border(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color("#7D56F4"))

	titleStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("#7D56F4")).
			MarginBottom(1)

	selectedStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#FFFFFF")).
			Background(lipgloss.Color("#7D56F4")).
			Padding(0, 1)

	normalStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#888888")).
			Padding(0, 1)

	completedStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#00FF00")).
			Strikethrough(true).
			Padding(0, 1)

	inputStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#FFFF00")).
			MarginTop(1)

	helpStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#666666")).
			MarginTop(1)
)

// TodoItem represents a single todo item.
type TodoItem struct {
	text      string
	completed bool
}

// Model represents the todo list panel state.
type Model struct {
	items     []TodoItem
	cursor    int
	inputMode bool
	input     string
	width     int
	height    int
}

// New creates a new todo list panel model with sample items.
func New() Model {
	return Model{
		items: []TodoItem{
			{text: "Write E2E tests for TUI app", completed: false},
			{text: "Set up CI/CD pipeline", completed: false},
			{text: "Review BubbleTea documentation", completed: true},
			{text: "Implement tab navigation", completed: true},
			{text: "Add keyboard shortcuts", completed: false},
		},
		cursor:    0,
		inputMode: false,
	}
}

// SetSize sets the panel dimensions.
func (m *Model) SetSize(width, height int) {
	m.width = width - 4
	m.height = height - 4
}

// Update handles messages for the todo list panel.
func (m Model) Update(msg tea.Msg) (Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if m.inputMode {
			switch msg.String() {
			case "enter":
				if m.input != "" {
					m.items = append(m.items, TodoItem{text: m.input, completed: false})
					m.input = ""
				}
				m.inputMode = false
			case "esc":
				m.input = ""
				m.inputMode = false
			case "backspace":
				if len(m.input) > 0 {
					m.input = m.input[:len(m.input)-1]
				}
			default:
				if len(msg.String()) == 1 {
					m.input += msg.String()
				}
			}
		} else {
			switch msg.String() {
			case "up", "k":
				if m.cursor > 0 {
					m.cursor--
				}
			case "down", "j":
				if m.cursor < len(m.items)-1 {
					m.cursor++
				}
			case "enter", " ":
				if len(m.items) > 0 {
					m.items[m.cursor].completed = !m.items[m.cursor].completed
				}
			case "a":
				m.inputMode = true
				m.input = ""
			case "d":
				if len(m.items) > 0 {
					m.items = append(m.items[:m.cursor], m.items[m.cursor+1:]...)
					if m.cursor >= len(m.items) && m.cursor > 0 {
						m.cursor--
					}
				}
			}
		}
	}

	return m, nil
}

// View renders the todo list panel.
func (m Model) View() string {
	title := titleStyle.Render("✅ Todo List Panel - Interactive Task Manager")

	var items string
	if len(m.items) == 0 {
		items = normalStyle.Render("No items. Press 'a' to add one.")
	} else {
		for i, item := range m.items {
			checkbox := "[ ]"
			if item.completed {
				checkbox = "[✓]"
			}

			line := checkbox + " " + item.text

			if i == m.cursor {
				items += selectedStyle.Render(line) + "\n"
			} else if item.completed {
				items += completedStyle.Render(line) + "\n"
			} else {
				items += normalStyle.Render(line) + "\n"
			}
		}
	}

	var inputArea string
	if m.inputMode {
		inputArea = inputStyle.Render("New item: " + m.input + "█")
	}

	help := helpStyle.Render("↑/↓: Navigate • Enter/Space: Toggle • a: Add • d: Delete • q: Quit")

	content := lipgloss.JoinVertical(lipgloss.Left, title, items, inputArea, help)

	return panelStyle.
		Width(m.width).
		Height(m.height).
		Render(content)
}

// GetItems returns the current todo items (useful for testing).
func (m Model) GetItems() []TodoItem {
	return m.items
}

// GetCompletedCount returns the number of completed items.
func (m Model) GetCompletedCount() int {
	count := 0
	for _, item := range m.items {
		if item.completed {
			count++
		}
	}
	return count
}
