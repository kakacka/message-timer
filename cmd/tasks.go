package cmd

import (
	"fmt"

	"github.com/kakacka/message-timer/internal/config"
	"github.com/kakacka/message-timer/internal/tasks"
	"github.com/spf13/cobra"
)

var (
	tasksFlags  tasks.Flags
	tasksDirCmd = &cobra.Command{
		Use:   "dir",
		Short: "Print programs config directory",
		Long:  "Print programs config directory",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("Config dir: %s\n", config.ProgramDir)
		},
	}
	tasksMakeCmd = &cobra.Command{
		Use:   "make", //WIP
		Short: "Configure a new task",
		Long:  "Configure a new task that can be later run or edited in config directory",
		Args:  cobra.ExactArgs(1),
	}
	tasksListCmd = &cobra.Command{
		Use:   "list", //WIP
		Short: "List tasks from config directory",
		Long:  "List tasks from config directory",
	}
	tasksCmd = &cobra.Command{
		Use:   "tasks",
		Short: "Has subcommands for running tasks",
		Long:  "Has subcommands for making and running pre-configurated tasks",
	}
)

func init() {
	tasksMakeCmd.Flags().StringVarP(&tasksFlags.InputCommand, "inputcommand", "i", "", "Command which gets executed and piped to task")
	tasksMakeCmd.Flags().StringVarP(&tasksFlags.OutputCommand, "outputcommand", "o", "", "Command which gets executed with piped task output")
	tasksMakeCmd.Flags().StringVarP(&tasksFlags.Separator, "separator", "s", ";", "Timestamp and message separator. Separates on first occurrence on a line.")
	tasksMakeCmd.Flags().StringVarP(&tasksFlags.File, "file", "f", "", "Input file")

	tasksCmd.AddCommand(tasksDirCmd)
	tasksCmd.AddCommand(tasksMakeCmd)
	tasksCmd.AddCommand(tasksListCmd)
	rootCmd.AddCommand(tasksCmd)
}
