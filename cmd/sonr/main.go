package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
)

func main() {
	log.Println("Will become bind executable later")

	// get `go` executable path
	goExecutable, _ := exec.LookPath("go")

	// construct `go version` command
	cmdGoVer := &exec.Cmd{
		Path:   goExecutable,
		Args:   []string{goExecutable, "version"},
		Stdout: os.Stdout,
		Stderr: os.Stdout,
	}

	// run `go version` command
	if err := cmdGoVer.Run(); err != nil {
		fmt.Println("Error:", err)
	}
}
