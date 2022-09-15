package proccess

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
	"time"

	"github.com/araddon/dateparse"
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
		panic(err)
	}
	if fi.Mode()&os.ModeNamedPipe == 0 && flags.File == "" {
		fmt.Println("Missing input data")
		return
	}

	var inputReader io.Reader
	if flags.File != "" {
		inputReader, err = os.Open(flags.File)
		if err != nil {
			fmt.Println(err)
			return
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
			fmt.Println(errx)
			return
		}
		linenum++
		var stimestamp string
		var message string
		if flags.Separator != "" {
			splitLine := strings.SplitN(line, flags.Separator, 2)
			if len(splitLine) != 2 {
				fmt.Printf("Couldn't separate on line %d\n", linenum)
				fmt.Print(line)
				return
			}
			stimestamp = splitLine[0]
			message = splitLine[1] + "\n"
		}
		timestamp, errx := decodeTimestamp(stimestamp, &flags)
		if errx != nil {
			fmt.Println(errx)
			return
		}
		if deltaSet == false {
			timeDelta = time.Now().Sub(timestamp)
			deltaSet = true

		} else {
			nextTime := timestamp.Add(timeDelta)
			for nextTime.After(time.Now()) {
				//drink cup of tea
			}
		}
		flags.Stdout.WriteString(message)
	}
	if errx := scanner.Err(); errx != nil {
		fmt.Println(errx)
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
