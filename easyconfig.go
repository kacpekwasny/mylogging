package mylogging

//////////////////////////////////////////////
// Easy config, use these functions for easy configuration

// there are other parameters that you might want to check out before starting
// but this will work just fine
// Remember though to LoggerInit(), *Logger.Listen() and *Logger.StopListen()

// RotatingFiles file handler easy config
// conOut - Output to console?
// dirpath - path of directory where logs will be saved (better make it empty, to avoid deleting
// 		other log files)
// max files - number of log files wont exceed this number, oldest will be deleted
// file size limit in bytes - after file exceeds this size new file is made, and oldest is deleted
// 		if to many files are present
func (l *Logger) RotatingFiles(conOut bool, dirpath string, maxfiles int, fileSizeLimitBytes int64) {
	l.SetOutToCon(conOut)
	l.WrToFile = true
	l.RotFileHand = true
	l.SetDirPath(dirpath)
	l.MaxFiles = maxfiles
	l.FileSizeLimit = fileSizeLimitBytes
}

// WrToSingleFile easy config of single file output log
// conOut - Output to console?
// filepath - path of file where the logs will be written to
// file size limit in bytes - after file exceeds this size it will be moved to filepath + ".old.log"
//		and new file at filepath is made where new logs will go
func (l *Logger) WrToSingleFile(conOut bool, filepath string, fileSizeLimitBytes int64) {
	l.SetOutToCon(conOut)
	l.WrToFile = true
	l.RotFileHand = false
	l.SetFilePath(filepath)
	l.FileSizeLimit = fileSizeLimitBytes
}

// OutToConsole easy config
// geting logs only out to console
func (l *Logger) OutToConsole() {
	l.OutToCon = true
	l.WrToFile = false
	l.RotFileHand = false
}

// LoggerInit initialize logger with default values
//
func LoggerInit() *Logger {
	l := Logger{
		WrToFile:    false,
		OutToCon:    true,
		RotFileHand: false,

		FilePerm:      0644,
		FileSizeLimit: 1024,
		MaxFiles:      10,
		Level:         5,
		MsgLengthLim:  200,
		Buff:          10,

		FileNameFormat: "%D%_%M%_%Y%-gologging",
		FormatStr:      "%L% %D%-%M%-%Y%; %h%:%m%:%s% >> ",
	}
	l.SetFileNameFormat(l.FileNameFormat)
	l.SetFormatStr(l.FormatStr)
	return &l
}
