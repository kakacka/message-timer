package cmd

import (
	proccess "github.com/kakacka/message-timer/internal/process"
	"github.com/spf13/cobra"
)

var (
	flags proccess.Flags

	runCmd = &cobra.Command{
		Use:   "run",
		Short: "run the program",
		Long:  `run the program...`,
		Run: func(cmd *cobra.Command, args []string) {
			proccess.Run(flags)
		},
	}
)

func init() {
	rootCmd.AddCommand(runCmd)

	runCmd.Flags().StringVarP(&flags.Separator, "separator", "s", ";", "Timestamp and message separator. Separates on first occurrence.")
	runCmd.Flags().StringVarP(&flags.File, "file", "f", "", "Input file")
}
