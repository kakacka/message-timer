package cmd

import (
	"os"

	"github.com/kakacka/message-timer/internal/config"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:     "message-timer",
	Short:   "Command line tool for timing messages",
	Long:    "Command line tool for timing messages with timestamps to simulate real-time dataflow. Each message has to be on each line in this format: (message)(separator)(timestamp). Messages are then written to stdout.",
	Version: config.Version,
}

func Execute() {
	config.LoadLogger()
	if errx := config.LoadConfigDir(); errx != nil {
		log.Error("Couldn't load config directory")
	}
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
}

func init() {
}
