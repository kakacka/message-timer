package config

import (
	"os"
	"path/filepath"

	log "github.com/sirupsen/logrus"
)

var Version = "1.1.0"
var ProgramName = "message-timer"
var ProgramDir = ""
var Debug = false

func LoadConfigDir() error {
	if configDir, errx := os.UserConfigDir(); errx != nil {
		return errx
	} else {
		ProgramDir = filepath.Join(configDir, ProgramName)
		if errx := os.MkdirAll(ProgramDir, os.ModePerm); errx != nil {
			return errx
		}
	}
	return nil
}

func LoadLogger() {
	log.SetFormatter(&log.TextFormatter{DisableLevelTruncation: true})
	if Debug {
		log.SetLevel(log.DebugLevel)
	}
}
