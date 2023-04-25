package serialisation

import (
	"errors"
	"testing"
)

func getTestCases(t *testing.T) map[string][]any {
	// map in format [encoded: decoded]
	return map[string][]any{
		"[a,b,c]": {"a", "b", "c"},
		// contains int, float, positive, negative, string, string started with i, string started with f, nested array
		"[i12,i13,abc,\\iabc,[\\fzxy,f22222.4,f-2222.5,f-12.2,i0]]": {12, 13, "abc", "iabc", []any{"fzxy", 22222.4, -2222.5, -12.2, 0}},
		"[\\\\\\\\]":                 {"\\\\"},          // esc byte
		"[\\\\\\\\,\\,\\,\\,\\,\\,]": {"\\\\", ",,,,,"}, // esc byte and sep byte
		"[\\\\\\\\,\\,\\,\\,\\,\\,,[\\[\\[,\\]\\]\\[\\[]]": {"\\\\", ",,,,,", []any{"[[", "]][["}}, // esc byte and sep byte and open arr byte and close array byte and nested array
	}
}

func TestEncode(t *testing.T) {
	cases := getTestCases(t)
	for expected, data := range cases {
		res := Encode(data)
		if res != expected {
			t.Errorf("Expected %s got %s", expected, res)
		}
	}
}

func TestDecodeOk(t *testing.T) {
	cases := getTestCases(t)
	for data, expected := range cases {
		res, err := Decode(data)
		if err != nil {
			t.Errorf("Expected no error, got %s", err)
			compareSlices(t, expected, res)
		}
	}
}

func compareSlices(t *testing.T, expected, actual []any) {
	if len(expected) != len(actual) {
		t.Errorf("Expected equal len, got %d and %d", len(expected), len(actual))
		return
	}
	for i := 0; i < len(expected); i++ {
		if expectedSlice, ok := expected[i].([]any); ok {
			if actualSlice, ok := actual[i].([]any); ok {
				compareSlices(t, expectedSlice, actualSlice)
				continue
			}
			t.Errorf("Expected both %d elements to be slices", i)
			continue
		}
		if expected[i] != actual[i] {
			t.Errorf("Expected both %d elements to be equal, got %v and %v", i, expected[i], actual[i])
		}
	}
}

func TestDecodeErrors(t *testing.T) {
	cases := map[string]error{
		"a[]":        ErrInvalidFormat, // not stating with array open
		"[[[[":       ErrInvalidFormat, // not valid array
		"[i123.23]":  ErrConversion,    // not valid integer
		"[i123bbzz]": ErrConversion,    // not valid integer
		"[f123bbzz]": ErrConversion,    // not valid float
		"[\\":        ErrInvalidFormat, // not closed array
		"[\\]":       ErrInvalidFormat, // not closed array after esc char
	}
	for encoded, expected := range cases {
		_, err := Decode(encoded)
		if !errors.Is(err, expected) {
			t.Errorf("Expected [%s] got [%s]", expected, err)
		}
	}
}
