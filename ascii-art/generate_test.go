package main

import (
	"strings"
	"testing"
)

func TestGenerateArt(t *testing.T) {
	// Mock banner for testing
	// 'A' and 'B' are 8 lines high
	mockBanner := map[rune][]string{
		'A': {" A ", "A A", "AAA", "A A", "A A", "A A", "A A", "A A"},
		'B': {"BB ", "B B", "BB ", "B B", "B B", "B B", "B B", "BB "},
	}

	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "Empty String",
			input:    "",
			expected: "",
		},
		{
			name:     "Single Newline",
			input:    "\\n",
			expected: "\n",
		},
		{
			name:     "Double Newline",
			input:    "\\n\\n",
			expected: "\n\n",
		},
		{
			name:  "Basic Word",
			input: "A",
			expected: strings.Join([]string{
				" A ", "A A", "AAA", "A A", "A A", "A A", "A A", "A A", "",
			}, "\n"),
		},
		{
			name:  "Two Words Separated by Newline",
			input: "A\\nB",
			expected: strings.Join([]string{
				" A ", "A A", "AAA", "A A", "A A", "A A", "A A", "A A", // A
				"BB ", "B B", "BB ", "B B", "B B", "B B", "B B", "BB ", // B
				"", // Final newline
			}, "\n"),
		},
		{
			name:  "Newline in Middle (Empty Segment)",
			input: "A\\n\\nB",
			expected: strings.Join([]string{
				" A ", "A A", "AAA", "A A", "A A", "A A", "A A", "A A", // A
				"",                                                     // The single \n from the empty segment
				"BB ", "B B", "BB ", "B B", "B B", "B B", "B B", "BB ", // B
				"", // Final newline
			}, "\n"),
		},
		{
			name:  "Leading and Trailing Newlines",
			input: "\\nA\\n",
			expected: strings.Join([]string{
				"",                                                     // Leading \n
				" A ", "A A", "AAA", "A A", "A A", "A A", "A A", "A A", // A
				"", // Final newline (trailing empty segment is usually ignored in piscine)
			}, "\n"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := GenerateArt(tt.input, mockBanner)
			if got != tt.expected {
				t.Errorf("\nFAIL: %s\nInput: %q\nEXPECTED:\n%q\nGOT:\n%q", tt.name, tt.input, tt.expected, got)
			}
		})
	}
}
