package stringutils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type stringTestCase struct {
	name           string
	input          string
	expectedEmpty  bool
	expectedString string
}

func TestIsEmpty(t *testing.T) {
	tests := []stringTestCase{
		{
			name:          "IsEmpty | empty string",
			input:         "",
			expectedEmpty: true,
		},
		{
			name:          "IsEmpty | empty whitespace",
			input:         "     ",
			expectedEmpty: true,
		},
		{
			name:          "IsEmpty | empty newlines",
			input:         "\n\n\n",
			expectedEmpty: true,
		},
		{
			name:          "IsEmpty | empty tabs",
			input:         "\n\t\t\n",
			expectedEmpty: true,
		},
		{
			name:          "IsEmpty | not empty",
			input:         "   hello there    ",
			expectedEmpty: false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, testIsEmpty(tc))
	}
}

func testIsEmpty(tc stringTestCase) func(*testing.T) {
	return func(t *testing.T) {
		isEmpty := IsEmpty(tc.input)
		assert.Equal(t, tc.expectedEmpty, isEmpty)
	}
}

func TestNotEmpty(t *testing.T) {
	tests := []stringTestCase{
		{
			name:          "NotEmpty | string",
			input:         "i am a string",
			expectedEmpty: false,
		},
		{
			name:          "NotEmpty | empty",
			input:         "",
			expectedEmpty: true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, testNotEmpty(tc))
	}
}

func testNotEmpty(tc stringTestCase) func(*testing.T) {
	return func(t *testing.T) {
		notEmpty := NotEmpty(tc.input)
		assert.Equal(t, !tc.expectedEmpty, notEmpty)
	}
}

func TestSquish(t *testing.T) {
	tests := []stringTestCase{
		{
			name:           "Squish - normal string",
			input:          "i am normal :)",
			expectedString: "i am normal :)",
		},
		{
			name:           "Squish - extra whitespace",
			input:          "     i am   spacious    ...",
			expectedString: "i am spacious ...",
		},
		{
			name:           "Squish - linebreaks",
			input:          "i\nhave\nmany\nlinebreaks",
			expectedString: "i have many linebreaks",
		},
		{
			name:           "Squish - whitespace and linebreaks",
			input:          "   i\n   am   \n  loaded  \n!",
			expectedString: "i am loaded !",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, testSquish(tc))
	}
}

func testSquish(tc stringTestCase) func(*testing.T) {
	return func(t *testing.T) {
		output := Squish(tc.input)
		assert.Equal(t, tc.expectedString, output)
	}
}
