package main

import (
	"log"

	"github.com/rithvikp/aviator"
	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use:   "aviator",
		Short: "Quickly and simply deploy applications",
		Run:   run,
	}
)

func main() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatalf("Unable to start aviator: %v\n", err)
	}
}

func run(cmd *cobra.Command, args []string) {
	a := aviator.Aviator{}
	err := a.Run()
	if err != nil {
		log.Fatalf("Unable to run aviator: %v\n", err)
	}
}
