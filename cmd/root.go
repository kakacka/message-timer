package cmd

import (
	"fmt"
	"os"

	"github.com/kakacka/message-timer/internal/config"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:     "message-timer",
	Short:   "Command line tool for timing messages",
	Long:    "Simple command line tool for timing messages based on timestamps",
	Version: config.Version,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func init() {
}
