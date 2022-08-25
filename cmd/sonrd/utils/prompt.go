package utils

import (
	"errors"
	"os"

	"github.com/jedib0t/go-pretty/table"
	"github.com/kataras/golog"
	"github.com/manifoldco/promptui"
	"github.com/sonr-io/sonr/pkg/motor"
)

func DisplayAcc(m motor.MotorNode, msg string) {
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.SetTitle("Account Information")
	t.AppendHeader(table.Row{"Address", "DID", "Balance"})
	t.AppendRows([]table.Row{
		{m.GetAddress(), m.GetDID().String(), m.GetBalance()},
	})
	t.SetStyle(table.StyleLight)
	t.Render()
	golog.Default.Printf("✅ SUCCESS: %s", msg)
}

func DisplayAccList(ul UserAuthList) {
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.SetTitle("Available Accounts")
	t.AppendHeader(table.Row{"#", "Address", "Password"})
	var rows []table.Row
	idx := 1
	for addr, ua := range ul.Auths {
		rows = append(rows, table.Row{idx, addr, ua.Password})
		idx++
	}
	t.AppendRows(rows)
	t.SetStyle(table.StyleLight)
	t.Render()
}

func PromptPassword() (string, error) {
	validate := func(input string) error {
		if len(input) < 8 {
			return errors.New("password must be at least 8 characters")
		}
		return nil
	}

	firstPrompt := promptui.Prompt{
		Mask:     rune('*'),
		Label:    "Password",
		Validate: validate,
	}
	initialPasswd, err := firstPrompt.Run()
	if err != nil {
		return "", err
	}

	secondPrompt := promptui.Prompt{
		Mask:     rune('*'),
		Label:    "Confirm Password",
		Validate: validate,
	}
	confirmPwd, err := secondPrompt.Run()
	if err != nil {
		return "", err
	}

	if initialPasswd == confirmPwd {
		return confirmPwd, nil
	}
	return "", errors.New("Incorrect password entered")
}
