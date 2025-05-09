package models

import (
	"time"
)

// TimeFormat is the standard time format used by Playtomic API
const TimeFormat = "2006-01-02T15:04:05"

// ParseTime parses a time string in Playtomic's format
func ParseTime(timeStr string) (time.Time, error) {
	return time.Parse(TimeFormat, timeStr)
}

// FormatTime formats a time.Time into Playtomic's expected format
func FormatTime(t time.Time) string {
	return t.Format(TimeFormat)
}
