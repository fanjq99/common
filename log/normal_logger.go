package log

import (
	"fmt"
	"log"
	"os"
)

type normalLogger struct {
	*log.Logger
	level int
}

func (l *normalLogger) SetLevel(lv int) {
	l.level = lv
}

func (l normalLogger) GetLevel() int {
	return l.level
}

func (l *normalLogger) Debug(v ...interface{}) {
	if l.level >= LevelDebug {
		l.Output(callDepth, header("DEBUG", fmt.Sprintln(v...)))
	}
}

func (l *normalLogger) Debugf(format string, v ...interface{}) {
	if l.level >= LevelDebug {
		l.Output(callDepth, header("DEBUG", fmt.Sprintf(format, v...)))
	}
}

func (l *normalLogger) Info(v ...interface{}) {
	if l.level >= LevelInfo {
		l.Output(callDepth, header("INFO", fmt.Sprintln(v...)))
	}
}

func (l *normalLogger) Infof(format string, v ...interface{}) {
	if l.level >= LevelInfo {
		l.Output(callDepth, header("INFO", fmt.Sprintf(format, v...)))
	}
}

func (l *normalLogger) Warn(v ...interface{}) {
	if l.level >= LevelWarn {
		l.Output(callDepth, header("WARN", fmt.Sprintln(v...)))
	}
}

func (l *normalLogger) Warnf(format string, v ...interface{}) {
	if l.level >= LevelWarn {
		l.Output(callDepth, header("WARN", fmt.Sprintf(format, v...)))
	}
}

func (l *normalLogger) Error(v ...interface{}) {
	if l.level >= LevelError {
		l.Output(callDepth, header("ERROR", fmt.Sprintln(v...)))
	}
}

func (l *normalLogger) Errorf(format string, v ...interface{}) {
	if l.level >= LevelError {
		l.Output(callDepth, header("ERROR", fmt.Sprintf(format, v...)))
	}
}

func (l *normalLogger) Fatal(v ...interface{}) {
	l.Output(callDepth, header("FATAL", fmt.Sprintln(v...)))
	os.Exit(1)
}

func (l *normalLogger) Fatalf(format string, v ...interface{}) {
	l.Output(callDepth, header("FATAL", fmt.Sprintf(format, v...)))
	os.Exit(1)
}

func (l *normalLogger) Panic(v ...interface{}) {
	l.Logger.Panic(v...)
}

func (l *normalLogger) Panicf(format string, v ...interface{}) {
	l.Logger.Panicf(format, v...)
}
