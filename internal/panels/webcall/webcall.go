// Package webcall provides a panel for making HTTP requests and displaying responses.
package webcall

import (
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"strings"
	"time"

	tea "github.com/charmbracelet/bubbletea/v2"
	"github.com/charmbracelet/lipgloss/v2"
)

var apiEndpoints = []struct {
	name string
	url  string
}{
	{"User #1", "https://jsonplaceholder.typicode.com/users/1"},
	{"User #2", "https://jsonplaceholder.typicode.com/users/2"},
	{"Post #1", "https://jsonplaceholder.typicode.com/posts/1"},
	{"Todo #1", "https://jsonplaceholder.typicode.com/todos/1"},
	{"Comment #1", "https://jsonplaceholder.typicode.com/comments/1"},
}

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

	disabledStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#555555")).
			Padding(0, 1)

	responseStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#00FF00"))

	errorStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#FF0000"))

	loadingBoxStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#FFFFFF")).
			Background(lipgloss.Color("#7D56F4")).
			Bold(true).
			Padding(1, 3).
			Align(lipgloss.Center)

	blockingOverlayStyle = lipgloss.NewStyle().
				Foreground(lipgloss.Color("#888888")).
				Italic(true)
)

// responseMsg carries the API response.
type responseMsg struct {
	data string
	err  error
}

// tickMsg is sent during loading animation.
type tickMsg time.Time

// Model represents the webcall panel state.
type Model struct {
	selected     int
	response     string
	loading      bool
	err          error
	width        int
	height       int
	spinnerFrame int
	loadingDots  int
}

// Spinner frames for loading animation.
var spinnerFrames = []string{"⠋", "⠙", "⠹", "⠸", "⠼", "⠴", "⠦", "⠧", "⠇", "⠏"}

// New creates a new webcall panel model.
func New() Model {
	return Model{
		selected: 0,
	}
}

// SetSize sets the panel dimensions.
func (m *Model) SetSize(width, height int) {
	m.width = width - 4
	m.height = height - 4
}

// getRandomDelay returns a random duration between 3-7 seconds.
func getRandomDelay() time.Duration {
	delay := 3000 + rand.Intn(4001)
	return time.Duration(delay) * time.Millisecond
}

// fetchData makes an HTTP request to the selected endpoint with a random delay.
func fetchData(url string) tea.Cmd {
	return func() tea.Msg {
		delay := getRandomDelay()
		time.Sleep(delay)

		client := &http.Client{Timeout: 15 * time.Second}
		resp, err := client.Get(url)
		if err != nil {
			return responseMsg{err: err}
		}
		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return responseMsg{err: err}
		}

		var prettyJSON map[string]interface{}
		if err := json.Unmarshal(body, &prettyJSON); err != nil {
			return responseMsg{data: string(body)}
		}

		formatted, err := json.MarshalIndent(prettyJSON, "", "  ")
		if err != nil {
			return responseMsg{data: string(body)}
		}

		return responseMsg{data: string(formatted)}
	}
}

// spinnerTick returns a command that sends a tick for the spinner animation.
func spinnerTick() tea.Cmd {
	return tea.Tick(100*time.Millisecond, func(t time.Time) tea.Msg {
		return tickMsg(t)
	})
}

// Update handles messages for the webcall panel.
func (m Model) Update(msg tea.Msg) (Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if m.loading {
			return m, nil
		}

		switch msg.String() {
		case "up", "k":
			if m.selected > 0 {
				m.selected--
			}
		case "down", "j":
			if m.selected < len(apiEndpoints)-1 {
				m.selected++
			}
		case "enter":
			m.loading = true
			m.response = ""
			m.err = nil
			m.spinnerFrame = 0
			m.loadingDots = 0
			return m, tea.Batch(fetchData(apiEndpoints[m.selected].url), spinnerTick())
		}

	case tickMsg:
		if m.loading {
			m.spinnerFrame = (m.spinnerFrame + 1) % len(spinnerFrames)
			m.loadingDots = (m.loadingDots + 1) % 4
			return m, spinnerTick()
		}

	case responseMsg:
		m.loading = false
		if msg.err != nil {
			m.err = msg.err
			m.response = ""
		} else {
			m.response = msg.data
			m.err = nil
		}
	}

	return m, nil
}

// truncateResponse limits the response to fit within available space.
func (m Model) truncateResponse(response string, maxLines int) string {
	lines := strings.Split(response, "\n")
	if len(lines) <= maxLines {
		return response
	}
	truncated := strings.Join(lines[:maxLines], "\n")
	return truncated + "\n... (truncated)"
}

// View renders the webcall panel.
func (m Model) View() string {
	title := titleStyle.Render("🌐 Web Call Panel - JSONPlaceholder API")

	// Render endpoint options
	var optionsBuilder strings.Builder
	if m.loading {
		optionsBuilder.WriteString(blockingOverlayStyle.Render("Select an endpoint (input disabled while loading):"))
	} else {
		optionsBuilder.WriteString("Select an endpoint (↑/↓) and press Enter to fetch:")
	}
	optionsBuilder.WriteString("\n\n")

	for i, endpoint := range apiEndpoints {
		if m.loading {
			marker := "  "
			if i == m.selected {
				marker = "▶ "
			}
			optionsBuilder.WriteString(disabledStyle.Render(marker + endpoint.name))
		} else {
			if i == m.selected {
				optionsBuilder.WriteString(selectedStyle.Render("▶ " + endpoint.name))
			} else {
				optionsBuilder.WriteString(normalStyle.Render("  " + endpoint.name))
			}
		}
		optionsBuilder.WriteString("\n")
	}

	options := optionsBuilder.String()

	// Render response area
	var responseArea string
	if m.loading {
		spinner := spinnerFrames[m.spinnerFrame]
		dots := strings.Repeat(".", m.loadingDots) + strings.Repeat(" ", 3-m.loadingDots)
		loadingMsg := fmt.Sprintf("%s Fetching data%s", spinner, dots)
		loadingBox := loadingBoxStyle.Render(loadingMsg)

		responseArea = "\n" + loadingBox + "\n\n"
		responseArea += blockingOverlayStyle.Render("⚠️  Please wait - request in progress (3-7 seconds)\n")
		responseArea += blockingOverlayStyle.Render("    Keyboard input is disabled until complete.")
	} else if m.err != nil {
		responseArea = errorStyle.Render(fmt.Sprintf("❌ Error: %v", m.err))
	} else if m.response != "" {
		// Calculate available lines for response (rough estimate)
		// Panel height minus title, options, and some padding
		availableLines := m.height - 12
		if availableLines < 5 {
			availableLines = 5
		}
		truncated := m.truncateResponse(m.response, availableLines)
		responseArea = responseStyle.Render("✅ Response:\n" + truncated)
	} else {
		responseArea = normalStyle.Render("Press Enter to fetch data from the selected endpoint.")
	}

	content := lipgloss.JoinVertical(lipgloss.Left, title, options, responseArea)

	// Force the panel to respect its height constraint
	return panelStyle.
		Width(m.width).
		Height(m.height).
		MaxHeight(m.height).
		Render(content)
}

// IsLoading returns whether the panel is currently loading (useful for testing).
func (m Model) IsLoading() bool {
	return m.loading
}

// GetResponse returns the current response (useful for testing).
func (m Model) GetResponse() string {
	return m.response
}

// GetSelectedEndpoint returns the currently selected endpoint name (useful for testing).
func (m Model) GetSelectedEndpoint() string {
	if m.selected >= 0 && m.selected < len(apiEndpoints) {
		return apiEndpoints[m.selected].name
	}
	return ""
}
