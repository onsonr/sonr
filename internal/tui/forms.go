package tui

import (
	"fmt"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/huh"
	"github.com/charmbracelet/lipgloss"
	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/tx"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	"github.com/spf13/cobra"
)

const maxWidth = 100

var (
	red    = lipgloss.AdaptiveColor{Light: "#FE5F86", Dark: "#FE5F86"}
	indigo = lipgloss.AdaptiveColor{Light: "#5A56E0", Dark: "#7571F9"}
	green  = lipgloss.AdaptiveColor{Light: "#02BA84", Dark: "#02BF87"}
)

type Styles struct {
	Base,
	HeaderText,
	Status,
	StatusHeader,
	Highlight,
	ErrorHeaderText,
	Help lipgloss.Style
}

func NewStyles(lg *lipgloss.Renderer) *Styles {
	s := Styles{}
	s.Base = lg.NewStyle().
		Padding(1, 2, 0, 1)
	s.HeaderText = lg.NewStyle().
		Foreground(indigo).
		Bold(true).
		Padding(0, 1, 0, 1)
	s.Status = lg.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(indigo).
		PaddingLeft(1).
		MarginTop(1)
	s.StatusHeader = lg.NewStyle().
		Foreground(green).
		Bold(true)
	s.Highlight = lg.NewStyle().
		Foreground(lipgloss.Color("212"))
	s.ErrorHeaderText = s.HeaderText.
		Foreground(red)
	s.Help = lg.NewStyle().
		Foreground(lipgloss.Color("240"))
	return &s
}

type state int

const (
	statusNormal state = iota
	stateDone
)

type Model struct {
	state   state
	lg      *lipgloss.Renderer
	styles  *Styles
	form    *huh.Form
	width   int
	message *tx.TxBody
}

func NewModel() Model {
	m := Model{width: maxWidth}
	m.lg = lipgloss.DefaultRenderer()
	m.styles = NewStyles(m.lg)

	m.form = huh.NewForm(
		huh.NewGroup(
			huh.NewInput().
				Key("from").
				Title("From Address").
				Placeholder("cosmos1...").
				Validate(func(s string) error {
					if !strings.HasPrefix(s, "cosmos1") {
						return fmt.Errorf("invalid address format")
					}
					return nil
				}),

			huh.NewInput().
				Key("to").
				Title("To Address").
				Placeholder("cosmos1...").
				Validate(func(s string) error {
					if !strings.HasPrefix(s, "cosmos1") {
						return fmt.Errorf("invalid address format")
					}
					return nil
				}),

			huh.NewInput().
				Key("amount").
				Title("Amount").
				Placeholder("100").
				Validate(func(s string) error {
					if _, err := sdk.ParseCoinNormalized(s + "atom"); err != nil {
						return fmt.Errorf("invalid coin amount")
					}
					return nil
				}),

			huh.NewSelect[string]().
				Key("denom").
				Title("Denom").
				Options(huh.NewOptions("atom", "osmo", "usnr", "snr")...),

			huh.NewInput().
				Key("memo").
				Title("Memo").
				Placeholder("Optional"),

			huh.NewConfirm().
				Key("done").
				Title("Ready to convert?").
				Validate(func(v bool) error {
					if !v {
						return fmt.Errorf("Please confirm when you're ready to convert")
					}
					return nil
				}).
				Affirmative("Yes, convert!").
				Negative("Not yet"),
		),
	).
		WithWidth(60).
		WithShowHelp(false).
		WithShowErrors(false)

	return m
}

func (m Model) Init() tea.Cmd {
	return m.form.Init()
}

func min(x, y int) int {
	if x > y {
		return y
	}
	return x
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = min(msg.Width, maxWidth) - m.styles.Base.GetHorizontalFrameSize()
	case tea.KeyMsg:
		switch msg.String() {
		case "esc", "ctrl+c", "q":
			return m, tea.Quit
		}
	}

	var cmds []tea.Cmd

	form, cmd := m.form.Update(msg)
	if f, ok := form.(*huh.Form); ok {
		m.form = f
		cmds = append(cmds, cmd)
	}

	if m.form.State == huh.StateCompleted {
		m.buildMessage()
		cmds = append(cmds, tea.Quit)
	}

	return m, tea.Batch(cmds...)
}

func (m Model) View() string {
	s := m.styles

	switch m.form.State {
	case huh.StateCompleted:
		pklCode := m.generatePkl()
		messageView := m.getMessageView()
		var b strings.Builder
		fmt.Fprintf(&b, "Final Tx:\n\n%s\n\n%s", pklCode, messageView)
		return s.Status.Margin(0, 1).Padding(1, 2).Width(80).Render(b.String()) + "\n\n"
	default:
		var schemaType string
		if m.form.GetString("schemaType") != "" {
			schemaType = "Schema Type: " + m.form.GetString("schemaType")
		}

		v := strings.TrimSuffix(m.form.View(), "\n\n")
		form := m.lg.NewStyle().Margin(1, 0).Render(v)

		var status string
		{
			preview := "(Preview will appear here)"
			if m.form.GetString("schema") != "" {
				preview = m.generatePkl()
			}

			const statusWidth = 40
			statusMarginLeft := m.width - statusWidth - lipgloss.Width(form) - s.Status.GetMarginRight()
			status = s.Status.
				Height(lipgloss.Height(form)).
				Width(statusWidth).
				MarginLeft(statusMarginLeft).
				Render(s.StatusHeader.Render("Pkl Preview") + "\n" +
					schemaType + "\n\n" +
					preview)
		}

		errors := m.form.Errors()
		header := m.appBoundaryView("Sonr TX Builder")
		if len(errors) > 0 {
			header = m.appErrorBoundaryView(m.errorView())
		}
		body := lipgloss.JoinHorizontal(lipgloss.Top, form, status)

		footer := m.appBoundaryView(m.form.Help().ShortHelpView(m.form.KeyBinds()))
		if len(errors) > 0 {
			footer = m.appErrorBoundaryView("")
		}

		return s.Base.Render(header + "\n" + body + "\n\n" + footer)
	}
}

func (m Model) errorView() string {
	var s string
	for _, err := range m.form.Errors() {
		s += err.Error()
	}
	return s
}

func (m Model) appBoundaryView(text string) string {
	return lipgloss.PlaceHorizontal(
		m.width,
		lipgloss.Left,
		m.styles.HeaderText.Render(text),
		lipgloss.WithWhitespaceChars("="),
		lipgloss.WithWhitespaceForeground(indigo),
	)
}

func (m Model) appErrorBoundaryView(text string) string {
	return lipgloss.PlaceHorizontal(
		m.width,
		lipgloss.Left,
		m.styles.ErrorHeaderText.Render(text),
		lipgloss.WithWhitespaceChars("="),
		lipgloss.WithWhitespaceForeground(red),
	)
}

func (m Model) generatePkl() string {
	schemaType := m.form.GetString("schemaType")
	schema := m.form.GetString("schema")

	// This is a placeholder for the actual conversion logic
	// In a real implementation, you would parse the schema and generate Pkl code
	return fmt.Sprintf("// Converted from %s\n\nclass ConvertedSchema {\n  // TODO: Implement conversion from %s\n  // Original schema:\n  /*\n%s\n  */\n}", schemaType, schemaType, schema)
}

func (m *Model) buildMessage() {
	from := m.form.GetString("from")
	to := m.form.GetString("to")
	amount := m.form.GetString("amount")
	denom := m.form.GetString("denom")
	memo := m.form.GetString("memo")

	coin, _ := sdk.ParseCoinNormalized(fmt.Sprintf("%s%s", amount, denom))
	sendMsg := &banktypes.MsgSend{
		FromAddress: from,
		ToAddress:   to,
		Amount:      sdk.NewCoins(coin),
	}

	anyMsg, _ := codectypes.NewAnyWithValue(sendMsg)
	m.message = &tx.TxBody{
		Messages: []*codectypes.Any{anyMsg},
		Memo:     memo,
	}
}

func (m Model) getMessageView() string {
	if m.message == nil {
		return "Current Message: None"
	}

	interfaceRegistry := codectypes.NewInterfaceRegistry()
	marshaler := codec.NewProtoCodec(interfaceRegistry)
	jsonBytes, _ := marshaler.MarshalJSON(m.message)

	return fmt.Sprintf("Current Message:\n%s", string(jsonBytes))
}

func RunTUIForm() (*tx.TxBody, error) {
	m := NewModel()
	p := tea.NewProgram(m)

	finalModel, err := p.Run()
	if err != nil {
		return nil, fmt.Errorf("failed to run program: %w", err)
	}

	finalM, ok := finalModel.(Model)
	if !ok || finalM.message == nil {
		return nil, fmt.Errorf("form not completed")
	}

	return finalM.message, nil
}

func NewTUIDashboardCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "dash",
		Short: "TUI for managing the local Sonr validator node",
		RunE: func(cmd *cobra.Command, args []string) error {
			txBody, err := RunTUIForm()
			if err != nil {
				return err
			}

			interfaceRegistry := codectypes.NewInterfaceRegistry()
			marshaler := codec.NewProtoCodec(interfaceRegistry)
			jsonBytes, err := marshaler.MarshalJSON(txBody)
			if err != nil {
				return fmt.Errorf("failed to marshal tx body: %w", err)
			}

			fmt.Println("Generated Protobuf Message (JSON format):")
			fmt.Println(string(jsonBytes))

			return nil
		},
	}
}
