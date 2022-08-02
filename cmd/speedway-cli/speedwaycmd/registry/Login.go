package MotorRegistry

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	mtr "github.com/sonr-io/sonr/pkg/motor"
	"github.com/spf13/cobra"
	rtmv1 "go.buf.build/grpc/go/sonr-io/motor/api/v1"
)

func loginAccount(did string, password string, pskKey []byte) {
	req, err := json.Marshal(rtmv1.LoginRequest{
		Did:       did,
		Password:  password,
		AesPskKey: pskKey,
	})
	if err != nil {
		fmt.Println("err", err)
	}
	fmt.Println("request", req)
	m := mtr.EmptyMotor("Speedway Node")
	res, err := m.Login(req)
	if err != nil {
		fmt.Println("err", err)
	}
	// if res returns false then the login failed
	fmt.Println("res", res)
	if res.Success {
		fmt.Println("Result", res)
		fmt.Println("DIDDocument", m.DIDDocument)
		fmt.Println("Address", m.Address)
		fmt.Println("Balance", m.Balance())
	} else {
		fmt.Println("Login failed")
	}
}

func loadKey(path string) ([]byte, error) {
	var file *os.File
	if _, err := os.Stat(path); err != nil {
		if os.IsNotExist(err) {
			return nil, err
		}
		return nil, err
	}
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	data, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	return data, nil
}

func bootstrapLoginCommand(ctx context.Context) (loginCmd *cobra.Command) {
	loginCmd = &cobra.Command{
		Use:   "loginAccount",
		Short: "Use: motor login [did] [password]",

		Run: func(cmd *cobra.Command, args []string) {
			if len(args) < 2 {
				fmt.Println("Please provide a did, password and pskKey")
				return
			}
			pskKey, err := loadKey("./PSK.key")
			if err != nil {
				fmt.Println("err", err)
			}
			did := args[0]
			password := args[1]
			loginAccount(did, password, pskKey)
		},
	}
	return
}
