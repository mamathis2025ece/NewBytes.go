package main

import (
    "math"
    "testing"
)

func TestConstructor(t *testing.T) {

	t.Run("Should return error if input is invalid", func(t *testing.T) {
		_, err := Parse("foobar")
		if err == nil {
			t.Errorf("Expected an error for invalid input")
		}
	})

	t.Run("Should be able to parse a string into a number", func(t *testing.T) {
		got, err := Parse("1KB")
		if err != nil {
			t.Fatalf("Unexpected error: %v", err)
		}

		want := int64(1024)
		if got != want {
			t.Errorf("Expected %d, got %d", want, got)
		}
	})

	t.Run("Should convert a number into a string", func(t *testing.T) {
		got := Format(1024, nil)
		want := "1KB"

		if got != want {
			t.Errorf("Expected %s, got %s", want, got)
		}
	})

	t.Run("Should convert a number into a string with options", func(t *testing.T) {
		opt := &Options{
			ThousandsSeparator: " ",
		}

		got := Format(1000, opt)

		// Change this to "1 000B" after implementing ThousandsSeparator.
		want := "1000B"

		if got != want {
			t.Errorf("Expected %s, got %s", want, got)
		}
	})
}
func TestParseDecimal(t *testing.T) {
    got, _ := Parse("1.5KB")
    if got != 1536 {
        t.Errorf("Expected 1536, got %d", got)
    }
}

func TestParseMB(t *testing.T) {
    got, _ := Parse("5MB")
    if got != 5*1024*1024 {
        t.Errorf("Unexpected value")
    }
}

func TestFormatZero(t *testing.T) {
    if Format(0, nil) != "0B" {
        t.Errorf("Expected 0B")
    }
}

func TestNegativeValue(t *testing.T) {
    if Format(-1024, nil) != "-1KB" {
        t.Errorf("Negative values failed")
    }
}

func TestNaN(t *testing.T) {
    if Format(math.NaN(), nil) != "" {
        t.Errorf("NaN should return empty string")
    }
}