package log

import (
	"fmt"
	"log"
	"os"

	"github.com/fatih/color"
)

type defaultLogger struct {
	*log.Logger
	level int
}

func (l *defaultLogger) SetLevel(lv int) {
	l.level = lv
}

func (l *defaultLogger) GetLevel() int {
	return l.level
}

func (l *defaultLogger) Debug(v ...interface{}) {
	if l.level >= LevelDebug {
		l.Output(callDepth, header("DEBUG", fmt.Sprintln(v...)))
	}
}

func (l *defaultLogger) Debugf(format string, v ...interface{}) {
	if l.level >= LevelDebug {
		l.Output(callDepth, header("DEBUG", fmt.Sprintf(format, v...)))
	}
}

func (l *defaultLogger) Info(v ...interface{}) {
	if l.level >= LevelInfo {
		l.Output(callDepth, header(color.GreenString("INFO "), fmt.Sprintln(v...)))
	}
}

func (l *defaultLogger) Infof(format string, v ...interface{}) {
	if l.level >= LevelInfo {
		l.Output(callDepth, header(color.GreenString("INFO "), fmt.Sprintf(format, v...)))
	}
}

func (l *defaultLogger) Warn(v ...interface{}) {
	if l.level >= LevelWarn {
		l.Output(callDepth, header(color.YellowString("WARN "), fmt.Sprintln(v...)))
	}
}

func (l *defaultLogger) Warnf(format string, v ...interface{}) {
	if l.level >= LevelWarn {
		l.Output(callDepth, header(color.YellowString("WARN "), fmt.Sprintf(format, v...)))
	}
}

func (l *defaultLogger) Error(v ...interface{}) {
	if l.level >= LevelError {
		l.Output(callDepth, header(color.RedString("ERROR"), fmt.Sprintln(v...)))
	}
}

func (l *defaultLogger) Errorf(format string, v ...interface{}) {
	if l.level >= LevelError {
		l.Output(callDepth, header(color.RedString("ERROR"), fmt.Sprintf(format, v...)))
	}
}

func (l *defaultLogger) Fatal(v ...interface{}) {
	l.Output(callDepth, header(color.MagentaString("FATAL"), fmt.Sprintln(v...)))
	os.Exit(1)
}

func (l *defaultLogger) Fatalf(format string, v ...interface{}) {
	l.Output(callDepth, header(color.MagentaString("FATAL"), fmt.Sprintf(format, v...)))
	os.Exit(1)
}

func (l *defaultLogger) Panic(v ...interface{}) {
	l.Logger.Panic(v...)
}

func (l *defaultLogger) Panicf(format string, v ...interface{}) {
	l.Logger.Panicf(format, v...)
}
