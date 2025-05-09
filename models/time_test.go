package models

import (
	"testing"
	"time"
)

func TestParseTime(t *testing.T) {
	timeStr := "2023-05-15T14:30:00"
	parsed, err := ParseTime(timeStr)
	if err != nil {
		t.Errorf("Expected successful parsing, got error: %v", err)
	}

	expected := time.Date(2023, 5, 15, 14, 30, 0, 0, time.UTC)
	if !parsed.Equal(expected) {
		t.Errorf("Expected %v, got %v", expected, parsed)
	}

	_, err = ParseTime("invalid-date")
	if err == nil {
		t.Errorf("Expected error for invalid date format, got no error")
	}
}

func TestFormatTime(t *testing.T) {
	testTime := time.Date(2023, 5, 15, 14, 30, 0, 0, time.UTC)
	formatted := FormatTime(testTime)

	expected := "2023-05-15T14:30:00"
	if formatted != expected {
		t.Errorf("Expected %q, got %q", expected, formatted)
	}
}
