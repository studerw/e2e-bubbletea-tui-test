// Package timer provides a stopwatch/timer panel.
package timer

import (
	"fmt"
	"time"

	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
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

	bigTimerStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("#FFFFFF")).
			Background(lipgloss.Color("#333333")).
			Padding(1, 4).
			Align(lipgloss.Center)

	buttonActiveStyle = lipgloss.NewStyle().
				Foreground(lipgloss.Color("#FFFFFF")).
				Background(lipgloss.Color("#7D56F4")).
				Padding(0, 2).
				Margin(0, 1)

	buttonInactiveStyle = lipgloss.NewStyle().
				Foreground(lipgloss.Color("#888888")).
				Background(lipgloss.Color("#333333")).
				Padding(0, 2).
				Margin(0, 1)

	lapStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#AAAAAA")).
			MarginTop(1)

	helpStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#666666")).
			MarginTop(1)
)

type tickMsg time.Time

// Model represents the timer panel state.
type Model struct {
	elapsed           time.Duration
	previouslyElapsed time.Duration
	startTime         time.Time
	running           bool
	laps              []time.Duration
	selected          int
	width             int
	height            int
}

// New creates a new timer panel model.
func New() Model {
	return Model{
		elapsed:           0,
		previouslyElapsed: 0,
		running:           false,
		laps:              make([]time.Duration, 0),
		selected:          0,
	}
}

// SetSize sets the panel dimensions.
func (m *Model) SetSize(width, height int) {
	m.width = width - 4
	m.height = height - 4
}

func tick() tea.Cmd {
	return tea.Tick(time.Millisecond*100, func(t time.Time) tea.Msg {
		return tickMsg(t)
	})
}

// Update handles messages for the timer panel.
func (m Model) Update(msg tea.Msg) (Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyPressMsg:
		switch msg.String() {
		case "left", "h":
			if m.selected > 0 {
				m.selected--
			}
		case "right", "l":
			if m.selected < 2 {
				m.selected++
			}
		case "enter", "space":
			switch m.selected {
			case 0:
				m.running = !m.running
				if m.running {
					m.startTime = time.Now()
					return m, tick()
				} else {
					m.previouslyElapsed = m.elapsed
				}
			case 1:
				if m.running {
					m.laps = append(m.laps, m.elapsed)
				}
			case 2:
				m.elapsed = 0
				m.previouslyElapsed = 0
				m.laps = make([]time.Duration, 0)
				m.running = false
			}
		case "s":
			m.running = !m.running
			if m.running {
				m.startTime = time.Now()
				return m, tick()
			} else {
				m.previouslyElapsed = m.elapsed
			}
		case "r":
			m.elapsed = 0
			m.previouslyElapsed = 0
			m.laps = make([]time.Duration, 0)
			m.running = false
		}

	case tickMsg:
		if m.running {
			m.elapsed = m.previouslyElapsed + time.Since(m.startTime)
			return m, tick()
		}
	}

	return m, nil
}

func formatDuration(d time.Duration) string {
	minutes := int(d.Minutes())
	seconds := int(d.Seconds()) % 60
	millis := int(d.Milliseconds()) % 1000 / 10
	return fmt.Sprintf("%02d:%02d.%02d", minutes, seconds, millis)
}

// View renders the timer panel.
func (m Model) View() string {
	title := titleStyle.Render("⏱️ Stopwatch Panel - Timer with Laps")

	timerDisplay := bigTimerStyle.Render(formatDuration(m.elapsed))

	var startBtn, lapBtn, resetBtn string
	startLabel := "▶ Start"
	if m.running {
		startLabel = "⏸ Stop"
	}

	if m.selected == 0 {
		startBtn = buttonActiveStyle.Render(startLabel)
	} else {
		startBtn = buttonInactiveStyle.Render(startLabel)
	}

	if m.selected == 1 {
		lapBtn = buttonActiveStyle.Render("🏁 Lap")
	} else {
		lapBtn = buttonInactiveStyle.Render("🏁 Lap")
	}

	if m.selected == 2 {
		resetBtn = buttonActiveStyle.Render("↺ Reset")
	} else {
		resetBtn = buttonInactiveStyle.Render("↺ Reset")
	}

	buttons := lipgloss.JoinHorizontal(lipgloss.Center, startBtn, lapBtn, resetBtn)

	var lapsDisplay string
	if len(m.laps) > 0 {
		lapsDisplay = "\nLaps:\n"
		for i, lap := range m.laps {
			lapsDisplay += lapStyle.Render(fmt.Sprintf("  Lap %d: %s", i+1, formatDuration(lap))) + "\n"
		}
	}

	help := helpStyle.Render("←/→: Select button • Enter/Space: Activate • s: Start/Stop • r: Reset")

	content := lipgloss.JoinVertical(
		lipgloss.Center,
		title,
		"\n",
		timerDisplay,
		"\n",
		buttons,
		lapsDisplay,
		"\n",
		help,
	)

	return panelStyle.
		Width(m.width).
		Height(m.height).
		Render(content)
}

// GetElapsed returns the elapsed duration (useful for testing).
func (m Model) GetElapsed() time.Duration {
	return m.elapsed
}

// IsRunning returns whether the timer is running (useful for testing).
func (m Model) IsRunning() bool {
	return m.running
}
