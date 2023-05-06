package tasks

import (
	"os"

	"github.com/kakacka/message-timer/internal/config"
	log "github.com/sirupsen/logrus"
)

type Flags struct {
	InputCommand  string
	OutputCommand string
	Separator     string
	File          string

	TimeFormat string
	Stdout     *os.File
	Stdin      *os.File
}

func Make(flags Flags) {
	if config.ProgramDir == "" {
		log.Fatal("Can't save because program directory isn't initiated")
	}

}
