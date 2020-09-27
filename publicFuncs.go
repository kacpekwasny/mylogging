package mylogging

import (
	"errors"
	"time"

	"github.com/kacpekwasny/pyfuncs"
)

// LogPrep d
func (l *Logger) LogPrep(level string, sl []string) error {
	var (
		parts         []string
		message       string = pyfuncs.Join(" ", sl)
		prefix        string
		lines         []string
		linesNoPrefix []string
	)

	for _, str := range l.formats {

		if str == "" {
			continue
		}

		if str == "D" {
			parts = append(parts, stri(time.Now().Day()))
		} else if str == "M" {
			parts = append(parts, stri(int(time.Now().Month())))
		} else if str == "Mon" {
			parts = append(parts, stri(time.Now().Month()))
		} else if str == "Y" {
			parts = append(parts, stri(time.Now().Year()))
		} else if str == "h" {
			parts = append(parts, stri(time.Now().Hour()))
		} else if str == "m" {
			parts = append(parts, stri(time.Now().Minute()))
		} else if str == "s" {
			parts = append(parts, stri(time.Now().Second()))
		} else if str == "L" {
			parts = append(parts, level)
		} else {
			parts = append(parts, str)
		}
	}
	prefix = pyfuncs.Join("", parts)

	if l.MsgLengthLim < 2 {
		panic(errors.New("length of msgLengthLimit has to be greater than 1"))
	}

	// Prepare lines
	// "\n" is added in function listenLoop
	splitNewLines := pyfuncs.Split(message, "\n")
	if len(splitNewLines) == 1 {
		linesNoPrefix = strToLines(splitNewLines[0], l.MsgLengthLim)
	} else {
		for _, s := range splitNewLines {
			if len(s) > l.MsgLengthLim {
				linesNoPrefix = append(linesNoPrefix, strToLines(s, l.MsgLengthLim)...)
			} else {
				linesNoPrefix = append(lines, s)
			}
		}
	}

	// Add prefixes
	for _, s := range linesNoPrefix {
		lines = append(lines, prefix+s)
	}

	return l.log(lines)
}

// Listen buff is the buffer for channel
func (l *Logger) Listen() {
	l.listenLoop(l.Buff)
}

// StopListen d
func (l *Logger) StopListen() {
	for len(l.LinesCh) > 0 {
		time.Sleep(time.Millisecond * 50)
	}
	l.File.Close()
}
