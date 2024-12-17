package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSum(t *testing.T) {
	t.Run("2 and 3, should return 5", func(t *testing.T) {
		result := Sum(2, 3)
		assert.Equal(t, 5, result)
	})
	t.Run("5 and 10, should return 15", func(t *testing.T) {
		result := Sum(5, 10)
		assert.Equal(t, 15, result)
	})
}

func TestSum2(t *testing.T) {
	testCase := []struct {
		name     string
		a        int
		b        int
		expected int
	}{
		{
			name:     "2 and 3, should return 5",
			a:        2,
			b:        3,
			expected: 5,
		},
		{
			name:     "5 and 10, should return 15",
			a:        5,
			b:        10,
			expected: 15,
		},
	}
	for _, tc := range testCase {
		t.Run(tc.name, func(t *testing.T) {
			result := Sum(tc.a, tc.b)
			assert.Equal(t, tc.expected, result)
		})
	}
}
