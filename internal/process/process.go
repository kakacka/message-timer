package proccess

import (
	"bufio"
	"io"
	"os"
	"strings"
	"time"

	"github.com/araddon/dateparse"
	log "github.com/sirupsen/logrus"
)

type Flags struct {
	Separator string
	File      string

	TimeFormat string
	Stdout     *os.File
	Stdin      *os.File
}

func Run(flags Flags) {
	fi, err := os.Stdin.Stat()
	if err != nil {
		log.Panic(err)
	}
	if fi.Mode()&os.ModeNamedPipe == 0 && flags.File == "" {
		log.Fatal("No data piped or file specified")
	}

	var inputReader io.Reader
	if flags.File != "" {
		inputReader, err = os.Open(flags.File)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		inputReader = flags.Stdin
	}
	scanner := bufio.NewScanner(inputReader)
	linenum := 0
	var timeDelta time.Duration
	deltaSet := false
	for scanner.Scan() {
		line := scanner.Text()
		if errx := scanner.Err(); errx != nil {
			log.Fatal(errx)
		}
		linenum++
		var stimestamp string
		var message string
		if flags.Separator != "" {
			splitLine := strings.SplitN(line, flags.Separator, 2)
			if len(splitLine) != 2 {
				log.Fatalf("Couldn't separate on line %d\n%s\n", linenum)
			}
			stimestamp = splitLine[0]
			message = splitLine[1] + "\n"
		}
		timestamp, errx := decodeTimestamp(stimestamp, &flags)
		if errx != nil {
			log.Fatal(errx)
		}
		if deltaSet == false {
			timeDelta = time.Now().Sub(timestamp)
			deltaSet = true

		} else {
			nextTime := timestamp.Add(timeDelta)
			for nextTime.After(time.Now()) {
				//waiting
			}
		}
		flags.Stdout.WriteString(message)
	}
	if errx := scanner.Err(); errx != nil {
		log.Error(errx)
	}
}

func decodeTimestamp(timestamp string, flags *Flags) (time.Time, error) {
	if flags.TimeFormat == "" {
		if timeFormat, errx := dateparse.ParseFormat(timestamp); errx != nil {
			return time.Time{}, errx

		} else {
			flags.TimeFormat = timeFormat
		}
	}
	DecodedTimestamp, errx := time.Parse(flags.TimeFormat, timestamp)
	if errx != nil { //try again incase format changed
		if timeFormat, errx := dateparse.ParseFormat(timestamp); errx != nil {
			return time.Time{}, errx

		} else {
			flags.TimeFormat = timeFormat
		}
		DecodedTimestamp, errx = time.Parse(flags.TimeFormat, timestamp)
	}
	return DecodedTimestamp, errx
}
