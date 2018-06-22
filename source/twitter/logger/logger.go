package logger

import (
	"gopkg.in/ChimeraCoder/anaconda.v2"
	"log"
)

func New(level uint) anaconda.Logger {
	return &logger{
		level: level,
	}
}

type logger struct {
	level uint
}

func (l *logger) Fatal(v ...interface{}) {
	if l.level <= LevelFatal {
		log.Fatal(v...)
	}
}

func (l *logger) Fatalf(s string, v ...interface{}) {
	if l.level <= LevelFatal {
		log.Fatalf(s, v...)
	}
}

func (l *logger) Panic(v ...interface{}) {
	if l.level <= LevelPanic {
		log.Panic(v...)
	}
}

func (l *logger) Panicf(s string, v ...interface{}) {
	if l.level <= LevelPanic {
		log.Panicf(s, v...)
	}
}

func (l *logger) Critical(v ...interface{}) {
	if l.level <= LevelCritical {
		log.Print(v...)
	}
}

func (l *logger) Criticalf(s string, v ...interface{}) {
	if l.level <= LevelCritical {
		log.Printf(s, v...)
	}
}

func (l *logger) Error(v ...interface{}) {
	if l.level <= LevelError {
		log.Print(v...)
	}
}

func (l *logger) Errorf(s string, v ...interface{}) {
	if l.level <= LevelError {
		log.Printf(s, v...)
	}
}

func (l *logger) Warning(v ...interface{}) {
	if l.level <= LevelWarning {
		log.Print(v...)
	}
}

func (l *logger) Warningf(s string, v ...interface{}) {
	if l.level <= LevelWarning {
		log.Printf(s, v...)
	}
}

func (l *logger) Notice(v ...interface{}) {
	if l.level <= LevelNotice {
		log.Print(v...)
	}
}

func (l *logger) Noticef(s string, v ...interface{}) {
	if l.level <= LevelNotice {
		log.Printf(s, v...)
	}
}

func (l *logger) Info(v ...interface{}) {
	if l.level <= LevelInfo {
		log.Print(v...)
	}
}

func (l *logger) Infof(s string, v ...interface{}) {
	if l.level <= LevelInfo {
		log.Printf(s, v...)
	}
}

func (l *logger) Debug(v ...interface{}) {
	if l.level <= LevelDebug {
		log.Print(v...)
	}
}

func (l *logger) Debugf(s string, v ...interface{}) {
	if l.level <= LevelDebug {
		log.Printf(s, v...)
	}
}
