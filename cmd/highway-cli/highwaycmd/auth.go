package highwaycmd

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"time"

	"github.com/duo-labs/webauthn/webauthn"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func bootstrapRegisterCommand(ctx context.Context) (registerCmd *cobra.Command) {
	// env
	var (
		RP_ORIGIN        = viper.GetString("RP_ORIGIN")
		RP_AUTH_ORIGIN   = viper.GetString("RP_AUTH_ORIGIN")
		HTTP_PORT        = viper.GetInt("HTTP_PORT")
		LOCAL_CONFIG_DIR = viper.GetString("LOCAL_CONFIG_DIR")
		API_ENDPOINT     = fmt.Sprintf("%s:%d/v1", RP_ORIGIN, HTTP_PORT)
	)

	registerCmd = &cobra.Command{
		Use:   "register <name>",
		Short: "Register a domain on the Sonr network",

		Run: func(cmd *cobra.Command, args []string) {
			if len(args) < 1 {
				cobra.CheckErr(errors.New("missing 'name' argument"))
			}

			fmt.Print("Registering... ")
			run(fmt.Sprintf(
				"%s?operation=register&rp_origin=%s&username=%s",
				RP_AUTH_ORIGIN,
				url.QueryEscape(API_ENDPOINT),
				url.QueryEscape(args[0])),
				LOCAL_CONFIG_DIR,
			)
		},
	}
	return
}

func bootstrapLoginCommand(ctx context.Context) (loginCmd *cobra.Command) {
	// env
	var (
		RP_ORIGIN        = viper.GetString("RP_ORIGIN")
		RP_AUTH_ORIGIN   = viper.GetString("RP_AUTH_ORIGIN")
		HTTP_PORT        = viper.GetInt("HTTP_PORT")
		LOCAL_CONFIG_DIR = viper.GetString("LOCAL_CONFIG_DIR")
		API_ENDPOINT     = fmt.Sprintf("%s:%d/v1", RP_ORIGIN, HTTP_PORT)
	)

	loginCmd = &cobra.Command{
		Use:   "login <name>",
		Short: "Login to the Sonr network",

		Run: func(cmd *cobra.Command, args []string) {
			if len(args) < 1 {
				cobra.CheckErr(errors.New("missing 'name' argument"))
			}

			fmt.Print("Logging in... ")
			run(fmt.Sprintf(
				"%s?operation=login&rp_origin=%s&username=%s",
				RP_AUTH_ORIGIN,
				url.QueryEscape(API_ENDPOINT),
				url.QueryEscape(args[0])),
				filepath.Join(LOCAL_CONFIG_DIR, "highway-cli", "login.json"),
			)
		},
	}
	return
}

func run(url, saveDir string) {
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
		if err := saveCredential(c, saveDir); err != nil {
			fmt.Printf("failed.")
			return
		}
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

func saveCredential(cred *webauthn.Credential, path string) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	b, err := json.Marshal(*cred)
	if err != nil {
		return err
	}
	_, err = file.Write(b)
	return err
}
