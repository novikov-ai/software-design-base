package average_calculator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_CalculateAverage(t *testing.T) {
	tests := []struct {
		name     string
		in       []int
		expected int
	}{
		{
			name:     "odd",
			in:       []int{1, 2, 3},
			expected: 2,
		},
		{
			name:     "even",
			in:       []int{1, 2},
			expected: 1,
		},
		{
			name:     "zero",
			in:       []int{-2, 2},
			expected: 0,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			calc := New()
			got := calc.CalculateAverage(tc.in)
			assert.Equal(t, tc.expected, got)
		})
	}
}
