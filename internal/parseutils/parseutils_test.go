package parseutils

import (
	"fmt"
	"testing"
	"time"

	"github.com/Prague-Kino/scraper/internal/errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type crownsTestCase struct {
	input    string
	expected int
	hasError bool
}

type dateTimeTestCase struct {
	name            string
	inputDate       time.Time
	inputTimeString string
	expected        time.Time
	expectedErr     error
}

func TestCrownsToInt(t *testing.T) {
	tests := []crownsTestCase{
		{
			input:    "190 Kč",
			expected: 190,
		},
		{
			input:    "100 CZK",
			expected: 100,
		},
		{
			input:    "1 cZk",
			expected: 1,
		},
		{
			input:    "1 czech crowns",
			expected: 0,
			hasError: true,
		},
	}

	for i, tc := range tests {
		name := fmt.Sprintf("CrownsToInt %d", i)
		t.Run(name, testCrownsToInt(tc))
	}
}

func testCrownsToInt(tc crownsTestCase) func(t *testing.T) {
	return func(t *testing.T) {
		output, err := CrownsToInt(tc.input)
		if tc.hasError {
			require.Error(t, err)
		} else {
			require.NoError(t, err)
		}

		assert.Equal(t, tc.expected, output)
	}
}

func TestCombineDateTime(t *testing.T) {
	tests := []dateTimeTestCase{
		{
			name: "CombineDateTime - success | date without time",
			inputDate: time.Date(
				2026, 2, 24, 0, 0, 0, 0, time.UTC,
			),
			inputTimeString: "22:30",
			expected: time.Date(
				2026, 2, 24, 22, 30, 0, 0, time.UTC,
			),
			expectedErr: nil,
		},
		{
			name: "CombineDateTime - success | date with time",
			inputDate: time.Date(
				2026, 7, 18, 12, 34, 56, 789, time.UTC,
			),
			inputTimeString: "11:00",
			expected: time.Date(
				2026, 7, 18, 11, 0, 0, 0, time.UTC,
			),
			expectedErr: nil,
		},
		{
			name:            "CombineDateTime - success | date with time",
			inputDate:       time.Now(),
			inputTimeString: "30 past seven",
			expectedErr: &errors.InvalidTimeFormatError{
				InvalidTime: "30 past seven",
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, testCombineDateTime(tc))
	}
}

func testCombineDateTime(tc dateTimeTestCase) func(t *testing.T) {
	return func(t *testing.T) {
		output, err := CombineDateTime(
			tc.inputDate,
			tc.inputTimeString,
		)

		if tc.expectedErr == nil {
			require.NoError(t, err)
			assert.Equal(t, tc.expected, output)
		} else {
			require.Error(t, err)
			assert.Equal(t, tc.expectedErr, err)
		}
	}
}
