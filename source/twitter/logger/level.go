package logger

// These constants are used by New to determine which print statements
// should be printed to the log.
const (
	LevelDebug = iota
	LevelInfo
	LevelNotice
	LevelWarning
	LevelError
	LevelCritical
	LevelPanic
	LevelFatal
)
