package highwaycmd

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os/exec"
	"runtime"
	"time"

	"github.com/duo-labs/webauthn/webauthn"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func bootstrapRegisterCommand(ctx context.Context) (registerCmd *cobra.Command) {
	// env
	var (
		RP_ORIGIN    = viper.GetString("RP_ORIGIN")
		HTTP_PORT    = viper.GetInt("HTTP_PORT")
		API_ENDPOINT = fmt.Sprintf("%s:%d/v1", RP_ORIGIN, HTTP_PORT)
	)

	registerCmd = &cobra.Command{
		Use:   "register <name>",
		Short: "Register a domain on the Sonr network",

		Run: func(cmd *cobra.Command, args []string) {
			if len(args) < 1 {
				cobra.CheckErr(errors.New("missing 'name' argument"))
			}

			fmt.Print("Registering... ")
			run(args[0], fmt.Sprintf(
				"http://localhost:8989?operation=register&rp_origin=%s&username=%s",
				url.QueryEscape(API_ENDPOINT),
				url.QueryEscape(args[0])),
			)
		},
	}
	return
}

func bootstrapLoginCommand(ctx context.Context) (loginCmd *cobra.Command) {
	// env
	var (
		RP_ORIGIN    = viper.GetString("RP_ORIGIN")
		HTTP_PORT    = viper.GetInt("HTTP_PORT")
		API_ENDPOINT = fmt.Sprintf("%s:%d/v1", RP_ORIGIN, HTTP_PORT)
	)

	loginCmd = &cobra.Command{
		Use:   "login <name>",
		Short: "Login to the Sonr network",

		Run: func(cmd *cobra.Command, args []string) {
			if len(args) < 1 {
				cobra.CheckErr(errors.New("missing 'name' argument"))
			}

			fmt.Print("Logging in... ")
			run(args[0], fmt.Sprintf(
				"http://localhost:8989?operation=login&rp_origin=%s&username=%s",
				url.QueryEscape(API_ENDPOINT),
				url.QueryEscape(args[0])),
			)
		},
	}
	return
}

func run(name, url string) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(time.Minute))
	defer cancel()

	sessChan := make(chan *webauthn.Credential)
	go startEphemeralAuthServer(sessChan, cancel)

	err := openBrowser(url)
	if err != nil {
		fmt.Println("Failed to open browser:", err)
		return
	}
	cobra.CheckErr(err)

	select {
	case <-ctx.Done():
		fmt.Println("failed.")
	case c := <-sessChan:
		fmt.Println("done.")
		fmt.Print("Saving session... ")
		fmt.Println(c.ID)
		fmt.Println("done.")
	}
}

func startEphemeralAuthServer(sessChan chan<- *webauthn.Credential, cancel context.CancelFunc) {
	http.Handle("/", http.FileServer(http.Dir("./cmd/highway-cli/static")))
	http.HandleFunc("/callback", func(w http.ResponseWriter, r *http.Request) {
		body, err := ioutil.ReadAll(r.Body)
		cobra.CheckErr(err)
		if err != nil {
			fmt.Println("Failed to read body:", err)
		}
		var session *webauthn.Credential
		cobra.CheckErr(json.Unmarshal(body, &session))
		sessChan <- session
	})
	http.HandleFunc("/error", func(w http.ResponseWriter, r *http.Request) {
		cancel()
	})
	cobra.CheckErr(http.ListenAndServe(":8989", nil))
}

func openBrowser(url string) error {
	var err error

	switch runtime.GOOS {
	case "linux":
		err = exec.Command("xdg-open", url).Start()
	case "windows":
		err = exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
	case "darwin":
		err = exec.Command("open", url).Start()
	default:
		err = fmt.Errorf("unsupported platform")
	}
	return err
}
