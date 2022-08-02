package MotorRegistry

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/sonr-io/sonr/pkg/crypto"
	mtr "github.com/sonr-io/sonr/pkg/motor"
	"github.com/spf13/cobra"
	rtmv1 "go.buf.build/grpc/go/sonr-io/motor/api/v1"
)

func storeKey(name string, key []byte) error {
	// TODO: use a better way to store keys
	file, err := os.Create(name)
	if err != nil {
		return err
	}
	defer file.Close()
	_, err = file.Write(key)
	if err != nil {
		return err
	}
	return nil
}

func bootstrapCreateAccountCommand(ctx context.Context) (createCmd *cobra.Command) {
	createCmd = &cobra.Command{
		Use:   "createAccount",
		Short: "Use: registry createAccount [password]",

		Run: func(cmd *cobra.Command, args []string) {
			if len(args) == 0 {
				fmt.Println("Please provide a password")
				return
			}
			password := args[0]
			aesKey, err := crypto.NewAesKey()
			if err != nil {
				fmt.Println("err", err)
			}
			storeKey("AES.key", aesKey)
			fmt.Println("aesKey", aesKey)
			req, err := json.Marshal(rtmv1.CreateAccountRequest{
				Password:  password,
				AesDscKey: aesKey,
			})
			fmt.Println("request", req)
			if err != nil {
				fmt.Println("reqBytes err", err)
			}
			m := mtr.EmptyMotor("Speedway Node")
			res, err := m.CreateAccount(req)
			if err != nil {
				fmt.Println("err", err)
			}
			fmt.Println("res", res)
			fmt.Println("PskKey", res.AesPsk)
			storeKey("PSK.key", res.AesPsk)
		},
	}
	return
}
