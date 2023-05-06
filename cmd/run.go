package cmd

import (
	"os"

	proccess "github.com/kakacka/message-timer/internal/process"
	"github.com/spf13/cobra"
)

var (
	runFlags proccess.Flags

	runCmd = &cobra.Command{
		Use:   "run",
		Short: "Runs the program",
		Long:  "Runs the program",
		Run: func(cmd *cobra.Command, args []string) {
			runFlags.Stdin = os.Stdin
			runFlags.Stdout = os.Stdout
			proccess.Run(runFlags)
		},
	}
)

func init() {
	rootCmd.AddCommand(runCmd)

	runCmd.Flags().StringVarP(&runFlags.Separator, "separator", "s", ";", "Timestamp and message separator. Separates on first occurrence for each line.")
	runCmd.Flags().StringVarP(&runFlags.File, "file", "f", "", "Set input file")
	runCmd.Flags().StringVarP(&runFlags.TimeFormat, "layout", "l", "", "Set specific datetime layout for decoding timestamps. Used in go pkg time.Parse. If none is specified, uses dateparse library to auto-detect format.")
}
