package logger

// Level is the log level threshold for for a Logger.
type Level uint8

// These constants are used by New to determine which print statements
// should be printed to the log.
const (
	LevelDebug Level = iota
	LevelInfo
	LevelNotice
	LevelWarning
	LevelError
	LevelCritical
	LevelPanic
	LevelFatal
)

// ParseLevel converts a string into a Level.
func ParseLevel(s string) Level {
	switch s {
	default:
		fallthrough
	case LevelDebug.String():
		return LevelDebug
	case LevelInfo.String():
		return LevelInfo
	case LevelNotice.String():
		return LevelNotice
	case LevelWarning.String():
		return LevelWarning
	case LevelError.String():
		return LevelError
	case LevelCritical.String():
		return LevelCritical
	case LevelPanic.String():
		return LevelPanic
	case LevelFatal.String():
		return LevelFatal
	}
}

// String returns the string representation of a Level type.
func (l Level) String() string {
	switch l {
	default:
		fallthrough
	case LevelDebug:
		return "debug"
	case LevelInfo:
		return "info"
	case LevelNotice:
		return "notice"
	case LevelWarning:
		return "warning"
	case LevelError:
		return "error"
	case LevelCritical:
		return "critical"
	case LevelPanic:
		return "panic"
	case LevelFatal:
		return "fatal"
	}
}
