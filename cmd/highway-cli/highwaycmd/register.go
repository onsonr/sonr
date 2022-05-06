package highwaycmd

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"os/exec"
	"runtime"

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
			fmt.Println(args)
			if len(args) < 1 {
				cobra.CheckErr(errors.New("missing 'name' argument"))
			}
			done := make(chan bool)
			go startEphemeralAuthServer(done)

			fmt.Printf("Registering name %s... ", args[0])
			err := openBrowser(fmt.Sprintf(
				"http://localhost:8989?operation=register&rp_origin=%s&username=%s",
				url.QueryEscape(API_ENDPOINT),
				url.QueryEscape(args[0])),
			)
			cobra.CheckErr(err)

			<-done
			fmt.Println("success")
		},
	}
	return
}

func startEphemeralAuthServer(done chan bool) {
	http.Handle("/", http.FileServer(http.Dir("./cmd/highway-cli/static")))
	http.HandleFunc("/callback", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("done.")
		fmt.Println(r.Body)
		done <- true
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
