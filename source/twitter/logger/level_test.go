package logger

import (
	"testing"
)

func TestParseLevel(t *testing.T) {
	if ParseLevel("debug") != LevelDebug {
		t.Fail()
	}
	if ParseLevel("info") != LevelInfo {
		t.Fail()
	}
	if ParseLevel("notice") != LevelNotice {
		t.Fail()
	}
	if ParseLevel("warning") != LevelWarning {
		t.Fail()
	}
	if ParseLevel("error") != LevelError {
		t.Fail()
	}
	if ParseLevel("critical") != LevelCritical {
		t.Fail()
	}
	if ParseLevel("panic") != LevelPanic {
		t.Fail()
	}
	if ParseLevel("fatal") != LevelFatal {
		t.Fail()
	}
	if ParseLevel("asdf") != LevelDebug {
		t.Fail()
	}
}
