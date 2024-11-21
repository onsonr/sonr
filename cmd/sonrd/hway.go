package main

import (
	"github.com/onsonr/sonr/pkg/gateway"
	"github.com/spf13/cobra"
)

func NewHwayCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "hway",
		Short: "Start the Sonr HWay",
		RunE:  runHway,
	}
}

func runHway(cmd *cobra.Command, args []string) error {
	e := gateway.New()
	return e.Start(":3000")
}
