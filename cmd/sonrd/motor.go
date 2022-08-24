package main

import (
	"fmt"
	"os"

	"github.com/manifoldco/promptui"
	"github.com/sonr-io/sonr/cmd/sonrd/utils"
	"github.com/sonr-io/sonr/pkg/motor"
	"github.com/sonr-io/sonr/third_party/types/common"
	mt "github.com/sonr-io/sonr/third_party/types/motor"
	"github.com/spf13/cobra"
)

func RootMotorCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "motor",
		Short: "Setup a local Motor instance",
	}
	cmd.AddCommand(loginCmd(), registerCmd())
	return cmd
}

func loginCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "login",
		Short: "Login to an existing sonr account on disk",
		Run: func(cmd *cobra.Command, args []string) {
			prompt := promptui.Prompt{
				Label: "Enter your Address",
			}
			accAddr, err := prompt.Run()
			if err != nil {
				logger.Errorf("Failed to run Prompt", err)
				return
			}
			ua, err := utils.GetUserAuth(accAddr)
			if err != nil {
				logger.Errorf("Failed to fetch UserAuth", err)
				return
			}
			req := mt.LoginRequest{
				Did:       accAddr,
				Password:  ua.Password,
				AesPskKey: ua.AesPSKKey,
				AesDscKey: ua.AesDSCKey,
			}
			m := setupMotor()
			_, err = m.Login(req)
			if err != nil {
				logger.Errorf("Failed to login with UserAuth", err)
				return
			}
			utils.DisplayAcc(m, "Logged In")
		},
	}
	return cmd
}

func registerCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "register",
		Short: "Create a new Sonr Account",
		Run: func(cmd *cobra.Command, args []string) {
			passwd, err := utils.PromptPassword()
			if err != nil {
				logger.Errorf("Failed to create account %e", err)
			}
			ua, err := utils.NewUserAuth(passwd)
			if err != nil {
				logger.Errorf("Error creating new AES Key", err)
				return
			}
			req, err := ua.GenAccountCreateRequest()
			if err != nil {
				logger.Errorf("Error creating new AES Key", err)
				return
			}
			m := setupMotor()
			res, err := m.CreateAccount(*req)
			if err != nil {
				fmt.Printf("CreateAccount Error: %s\n", err)
				return
			}
			if err := ua.StoreAuth(res.Address, res.GetAesPsk()); err != nil {
				logger.Errorf("Failed to save UserAuth to Keychain", err)
				return
			}
			utils.DisplayAcc(m, "Account Registered")
		},
	}
	return cmd
}

func setupMotor() motor.MotorNode {
	initreq := &mt.InitializeRequest{
		DeviceId: utils.DesktopID(),
	}
	m, err := motor.EmptyMotor(initreq, common.DefaultCallback())
	if err != nil {
		fmt.Println("[FATAL] Error initializing motor node")
		os.Exit(1)
	}
	logger.Info("Motor Node has been initialized")
	return m
}
