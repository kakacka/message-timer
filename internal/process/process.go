package proccess

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
	"time"
)

type Flags struct {
	Separator string
	File      string
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
		inputReader = os.Stdin
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
		var value string
		if flags.Separator != "" {
			splitLine := strings.SplitN(line, flags.Separator, 2)
			if len(splitLine) != 2 {
				fmt.Printf("Couldn't separate on line %d\n", linenum)
				fmt.Print(line)
				return
			}
			stimestamp = splitLine[0]
			value = splitLine[1] + "\n"
		}
		timestamp, errx := formatTimestamp(stimestamp, flags)
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
		os.Stdout.WriteString(value)
	}
	if errx := scanner.Err(); errx != nil {
		fmt.Println(errx)
	}
}

func formatTimestamp(timestamp string, flags Flags) (time.Time, error) {
	return time.Parse("2006-01-02T15:04:05.999Z", timestamp)
}
