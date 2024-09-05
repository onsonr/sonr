package tui

import (
	"fmt"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/huh"
	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/tx"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	"github.com/spf13/cobra"
)

type Model struct {
	form    *huh.Form
	message *tx.TxBody
}

func NewModel() Model {
	return Model{
		form: huh.NewForm(
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
			),
		),
	}
}

func (m Model) Init() tea.Cmd {
	return m.form.Init()
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	form, cmd := m.form.Update(msg)
	if f, ok := form.(*huh.Form); ok {
		m.form = f
		if m.form.State == huh.StateCompleted {
			m.buildMessage()
		}
	}
	return m, cmd
}

func (m Model) View() string {
	formView := m.form.View()
	messageView := m.getMessageView()
	return fmt.Sprintf("%s\n\n%s", formView, messageView)
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

func runTUIForm() (*tx.TxBody, error) {
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

func NewBuildProtoMsgCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "build-proto-msg",
		Short: "Build a Cosmos SDK protobuf message using a TUI form",
		RunE: func(cmd *cobra.Command, args []string) error {
			txBody, err := runTUIForm()
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
