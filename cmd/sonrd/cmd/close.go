/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"io/ioutil"
	"os"
	"strconv"

	"github.com/kataras/golog"
	"github.com/sonr-io/core/app"
	"github.com/spf13/cobra"
)

var isErrorPtr bool

// closeCmd represents the close command
var closeCmd = &cobra.Command{
	Use:   "close",
	Short: "Closes any active Sonr Daemon thats running",
	Run: func(cmd *cobra.Command, args []string) {
		data, err := ioutil.ReadFile(PIDFile)
		if err != nil {
			golog.Child("[Sonr.daemon] ").Error("Not running")
			app.Exit(1)
		}

		ProcessID, err := strconv.Atoi(string(data))
		if err != nil {
			golog.Child("[Sonr.daemon] ").Error("Unable to read and parse process id found in ", PIDFile)
			app.Exit(1)
		}

		process, err := os.FindProcess(ProcessID)
		if err != nil {
			golog.Child("[Sonr.daemon] ").Errorf("Unable to find process ID [%v] with error %v \n", ProcessID, err)
			app.Exit(1)
		}
		// remove PID file
		os.Remove(PIDFile)
		err = process.Kill()
		if err != nil {
			golog.Child("[Sonr.daemon] ").Errorf("Unable to kill process ID [%v] with error %v \n", ProcessID, err)
			app.Exit(1)
		}
		app.Exit(0)
	},
}

func init() {
	rootCmd.AddCommand(closeCmd)
	closeCmd.Flags().BoolVar(&isErrorPtr, "error", false, "exits app with error code")
}
