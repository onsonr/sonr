package main

import (
	_ "embed"
	"fmt"
	"os"
)

// main is the entry point for the application
func main() {
	cmd := NewRootCmd()
	if err := cmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	os.Exit(0)
}
