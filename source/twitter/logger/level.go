package logger

// These constants are used by New to determine which print statements
// should be printed to the log.
const (
	LevelDebug uint = iota
	LevelInfo
	LevelNotice
	LevelWarning
	LevelError
	LevelCritical
	LevelPanic
	LevelFatal
)

// These constants are the string version of each log level. ParseLevel
// will translate these strings to the respective uint.
const (
	LevelDebugString    = "debug"
	LevelInfoString     = "info"
	LevelNoticeString   = "notice"
	LevelWarningString  = "warning"
	LevelErrorString    = "error"
	LevelCriticalString = "critical"
	LevelPanicString    = "panic"
	LevelFatalString    = "fatal"
)

// ParseLevel converts a string to a Level* uint.
func ParseLevel(s string) uint {
	switch s {
	default:
		fallthrough
	case LevelDebugString:
		return LevelDebug
	case LevelInfoString:
		return LevelInfo
	case LevelNoticeString:
		return LevelNotice
	case LevelWarningString:
		return LevelWarning
	case LevelErrorString:
		return LevelError
	case LevelCriticalString:
		return LevelCritical
	case LevelPanicString:
		return LevelPanic
	case LevelFatalString:
		return LevelFatal
	}
}
