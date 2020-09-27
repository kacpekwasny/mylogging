package mylogging

import cmt "github.com/kacpekwasny/commontools"

// Debug level 10
func (l *Logger) Debug(sl ...string) {
	if l.Level <= 10 {
		cmt.A(l.LogPrep("DEBUG", sl))
	}

}

// Info level 20
func (l *Logger) Info(sl ...string) {
	if l.Level <= 20 {
		cmt.A(l.LogPrep("INFO", sl))
	}
}

// Warning level 30
func (l *Logger) Warning(sl ...string) {
	if l.Level <= 30 {
		cmt.A(l.LogPrep("WARNING", sl))
	}
}

// Error level 40
func (l *Logger) Error(sl ...string) {
	if l.Level <= 40 {
		cmt.A(l.LogPrep("ERROR", sl))
	}
}

// Critical level 50
func (l *Logger) Critical(sl ...string) {
	if l.Level <= 50 {
		cmt.A(l.LogPrep("CRITICAL", sl))
	}
}
