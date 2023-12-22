package main

import (
	"errors"
	"testing"
)

func TestUnpackString(t *testing.T) {
	tests := []struct {
		input          string
		expectedOutput string
	}{
		{"a2b3c4", "aabbbcccc"},
		{"a0b1c2", "a0bcc"},
		{"a5", "aaaaa"},
		{"3a2b1c", "3aabc"},
		{"", ""},
		{"abc123def456", "abc222deffff555555"},
	}
	for _, test := range tests {
		t.Run(test.input, func(t *testing.T) {
			result := unpackString(test.input)
			if result != test.expectedOutput {
				t.Errorf("%s not equals expected %s \n", result, test.expectedOutput)
			}
		})

	}
}

func TestUnpackEscape(t *testing.T) {
	testCases := []struct {
		input    string
		expected string
		err      error
	}{
		{"abc", "abc", nil},
		{"a2b3c4", "aabbccc", nil},
		{"a0b1c2", "ac", nil},
		{"a5", "aaaaa", nil},
		{"3a2b1c", "aaabbc", nil},
		{"\\3a2b1c", "333a2b1c", nil},
		{"\\3a\\2b\\1c", "333aabbc", nil},
		{"\\a", "", errors.New("invalid string")},
		{"\\", "", errors.New("invalid string")},
		{"", "", errors.New("invalid string")},
	}

	for _, tc := range testCases {
		t.Run(tc.input, func(t *testing.T) {
			result, err := unpackEscape(tc.input)
			if err != nil && tc.err != nil {
				if err.Error() != tc.err.Error() {
					t.Errorf("Expected error '%v', but got '%v' for input '%s'", tc.err, err, tc.input)
				}

			} else if err != nil && tc.err == nil {
				t.Errorf("Expected error '%v', but got '%v' for input '%s'", tc.err, err, tc.input)

			} else if err == nil && tc.err != nil {
				t.Errorf("Expected error '%v', but got '%v' for input '%s'", tc.err, err, tc.input)
			} else {
				if result != tc.expected {
					t.Errorf("Expected '%s', but got '%s' for input '%s'", tc.expected, result, tc.input)
				}
			}
		})
	}
}
