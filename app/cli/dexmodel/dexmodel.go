package dexmodel

import (
	"fmt"
	"time"

	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/spf13/cobra"
)

var (
	subtle    = lipgloss.AdaptiveColor{Light: "#D9DCCF", Dark: "#383838"}
	highlight = lipgloss.AdaptiveColor{Light: "#874BFD", Dark: "#7D56F4"}
	special   = lipgloss.AdaptiveColor{Light: "#43BF6D", Dark: "#73F59F"}

	titleStyle = lipgloss.NewStyle().
			MarginLeft(1).
			MarginRight(5).
			Padding(0, 1).
			Italic(true).
			Foreground(lipgloss.Color("#FFF7DB")).
			SetString("Cosmos Block Explorer")

	infoStyle = lipgloss.NewStyle().
			BorderStyle(lipgloss.NormalBorder()).
			BorderTop(true).
			BorderForeground(subtle)
)

type model struct {
	blocks           []string
	transactionTable table.Model
	stats            map[string]string
	width            int
	height           int
}

func initialModel() model {
	columns := []table.Column{
		{Title: "Hash", Width: 10},
		{Title: "Type", Width: 15},
		{Title: "Height", Width: 10},
		{Title: "Time", Width: 20},
	}

	rows := []table.Row{
		{"abc123", "Transfer", "1000", time.Now().Format(time.RFC3339)},
		{"def456", "Delegate", "999", time.Now().Add(-1 * time.Minute).Format(time.RFC3339)},
		{"ghi789", "Vote", "998", time.Now().Add(-2 * time.Minute).Format(time.RFC3339)},
	}

	t := table.New(
		table.WithColumns(columns),
		table.WithRows(rows),
		table.WithFocused(true),
		table.WithHeight(7),
	)

	s := table.DefaultStyles()
	s.Header = s.Header.
		BorderStyle(lipgloss.NormalBorder()).
		BorderForeground(lipgloss.Color("240")).
		BorderBottom(true).
		Bold(false)
	s.Selected = s.Selected.
		Foreground(lipgloss.Color("229")).
		Background(lipgloss.Color("57")).
		Bold(false)
	t.SetStyles(s)

	return model{
		blocks:           []string{"Block 1", "Block 2", "Block 3"},
		transactionTable: t,
		stats: map[string]string{
			"Latest Block":  "1000",
			"Validators":    "100",
			"Bonded Tokens": "1,000,000",
		},
	}
}

func (m model) Init() tea.Cmd {
	return tick
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+c":
			return m, tea.Quit
		case "enter":
			return m, tea.Batch(
				tea.Printf("Selected transaction: %s", m.transactionTable.SelectedRow()[0]),
			)
		}
	case tea.WindowSizeMsg:
		m.height = msg.Height
		m.width = msg.Width
	case tickMsg:
		// Update data here
		m.blocks = append([]string{"New Block"}, m.blocks...)
		if len(m.blocks) > 5 {
			m.blocks = m.blocks[:5]
		}

		// Add a new transaction to the table
		newRow := table.Row{
			fmt.Sprintf("tx%d", time.Now().Unix()),
			"NewTxType",
			fmt.Sprintf("%d", 1000+len(m.transactionTable.Rows())),
			time.Now().Format(time.RFC3339),
		}
		m.transactionTable.SetRows(append([]table.Row{newRow}, m.transactionTable.Rows()...))
		if len(m.transactionTable.Rows()) > 10 {
			m.transactionTable.SetRows(m.transactionTable.Rows()[:10])
		}

		return m, tick
	}
	m.transactionTable, cmd = m.transactionTable.Update(msg)
	return m, cmd
}

func (m model) View() string {
	s := titleStyle.Render("Cosmos Block Explorer")
	s += "\n\n"

	// Blocks
	s += lipgloss.NewStyle().Bold(true).Render("Recent Blocks") + "\n"
	for _, block := range m.blocks {
		s += "• " + block + "\n"
	}
	s += "\n"

	// Transactions
	s += lipgloss.NewStyle().Bold(true).Render("Recent Transactions") + "\n"
	s += m.transactionTable.View() + "\n\n"

	// Stats
	s += lipgloss.NewStyle().Bold(true).Render("Network Statistics") + "\n"
	for key, value := range m.stats {
		s += fmt.Sprintf("%s: %s\n", key, value)
	}

	return s
}

type tickMsg time.Time

func tick() tea.Msg {
	time.Sleep(time.Second)
	return tickMsg{}
}

func RunExplorerTUI(cmd *cobra.Command, args []string) error {
	p := tea.NewProgram(initialModel(), tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		return fmt.Errorf("error running explorer: %v", err)
	}
	return nil
}
