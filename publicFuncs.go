package mylogging

import (
	"errors"
	"imports/pyfuncs"
	"time"
)

// LogPrep d
func (l *Logger) LogPrep(level string, sl []string) error {
	var (
		parts   []string
		message string = pyfuncs.Join(" ", sl)
		prefix  string
		lines   []string
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
	if len(message) > l.MsgLengthLim {
		var i int
		for j := range pyfuncs.Range(float64(int(len(message)/l.MsgLengthLim)) - 1) {
			i = int(j)
			lines = append(lines, prefix+message[i*l.MsgLengthLim:(i+1)*l.MsgLengthLim])
		}
		i++
		lines = append(lines, prefix+message[i*l.MsgLengthLim:])
	} else {
		lines = append(lines, prefix+message)
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
