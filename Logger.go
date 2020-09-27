package mylogging

import (
	"errors"
	"os"

	"github.com/kacpekwasny/pyfuncs"
)

// Logger keeps info on logging configuration
type Logger struct {
	// fileNameFormat format for roatating files
	// look at 'format' commments but no message
	// file name will be ended with .%nr%.log
	// and this is how files will be recognized
	FileNameFormat string

	// When you want multiple smaller files - rotating files
	// if rotating files then specify path to dir where logs will be saved
	DirPath string

	// when size of file exceeds the size limit next file is created
	FileSizeLimit int64
	FileSize      int64
	MaxFiles      int

	// if single log file then write to filePath
	FilePath string

	// permissions for log files
	FilePerm os.FileMode

	// *os.File object to which current output is written to
	File *os.File

	// Write to file/files ?
	// Out to console ?
	// rotating File Handler ? or single file
	WrToFile    bool
	OutToCon    bool
	RotFileHand bool

	// string that contains info on how output should be formatted
	//
	// DO NOT PUT % except for setting the format prefixes 	!!!!!!!!!!!!!!!!!!!!!!!!!!
	//
	// date and time
	// D-day, M-month number, Y-year, h-hour, m-minute, s-second, Mon-Month name
	// msg - message, L-debug levelName (INFO, ERROR...)
	// example:
	// %D%-%M% | %h%:%m% ->
	// and OUTPUT of this format string
	// '21-3 | 19:46 -> hello world'
	FormatStr string // for faster formatting purposes
	formats   []string

	// When the limit is exceeded the message will be wraped to next line
	// example, limit is 12 and message 012345678901234567890123456789:
	// '21-3 | 19:46 -> 012345678901'
	// '21-3 | 19:46 -> 234567890123'
	// '21-3 | 19:46 -> 456789'
	MsgLengthLim int

	// level is for depth of debugging
	Level int

	LinesCh chan string

	// Channel above buffer
	Buff int
}

// OUTPUT OPTIONS

// SetOutToCon f
func (l *Logger) SetOutToCon(isOn bool) {
	l.OutToCon = isOn
}

// SetWrToFile f
func (l *Logger) SetWrToFile(isOn bool) {
	l.WrToFile = isOn
}

// SetRotFileHand f
func (l *Logger) SetRotFileHand(isOn bool) {
	l.RotFileHand = isOn
	if isOn {
		l.WrToFile = true
	}
}

// SetFileNameFormat d
func (l *Logger) SetFileNameFormat(formatStr string) {
	l.FileNameFormat = formatStr
}

// SetFormatStr read Logger.FormatStr for info
func (l *Logger) SetFormatStr(formatStr string) {
	l.FormatStr = formatStr
	l.formats = []string{}
	frmts := pyfuncs.Split(formatStr, "%")
	for _, s := range frmts {
		if s != "" {
			l.formats = append(l.formats, s)
		}
	}
}

// SetDirPath dir path for log files in
func (l *Logger) SetDirPath(path string) {
	if string(path[len(path)-1]) != "/" {
		l.DirPath = path + "/"
	} else {
		l.DirPath = path
	}
}

// SetFileSizeLimit after exceeding new file is made
func (l *Logger) SetFileSizeLimit(size int64) {
	l.FileSizeLimit = size
}

// SetFilePath file path for log file and old log file with filepath + ".old.log"
func (l *Logger) SetFilePath(path string) {
	l.FilePath = path
}

// SetBuff for channel communication to logging goroutine
func (l *Logger) SetBuff(b int) {
	if b < 0 {
		panic(errors.New("buffer cannot be smaller than 0"))
	}
}

// SetMsgLengthLim Message Length after which exceeding text will be wrapped to next line
func (l *Logger) SetMsgLengthLim(lim int) {
	if lim > 0 {
		l.MsgLengthLim = lim
	} else {
		panic(errors.New("length of message cannot be smaller than 1"))
	}
}

// SetLevel level of the logger
func (l *Logger) SetLevel(v int) {
	l.Level = v
}
