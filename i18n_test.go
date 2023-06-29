package main

import (
	"testing"
)

type AcceptLanguageTest struct {
	AcceptLanguageHeader string
	ExpectedResult       string
}

func TestGetLanguageCodeFromHeader(t *testing.T) {
	var acceptLanguageTests = []AcceptLanguageTest{
		{"", "en"},
		{"en", "en"},
		{"es", "es"},
		{"en-US, fr;q=0.8, es;q=0.5", "en"},
		{"es-ES,es;q=0.9,en;q=0.8", "es"},
		{"en-GB", "en"},
		{"bu", "en"},
	}

	for _, test := range acceptLanguageTests {
		result := getLanguageCodeFromHeader(test.AcceptLanguageHeader)
		if result != test.ExpectedResult {
			t.Errorf("Error. Expected: %s, but get: %s with code: %s", test.ExpectedResult, result, test.AcceptLanguageHeader)
		}
	}
}
