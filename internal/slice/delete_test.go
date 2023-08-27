package slice

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/flipped94/gokit/internal/err"
)

func TestDeleteAt(t *testing.T) {
	testCases := []struct {
		name          string
		slice         []int
		index         int
		expectedSlice []int
		expectedVal   int
		expectedError error
	}{
		{
			name:          "index 0",
			slice:         []int{123, 100},
			index:         0,
			expectedSlice: []int{100},
			expectedVal:   123,
		},
		{
			name:          "index middle",
			slice:         []int{123, 124, 125},
			index:         1,
			expectedSlice: []int{123, 125},
			expectedVal:   124,
		},
		{
			name:          "index out of range",
			slice:         []int{123, 100},
			index:         12,
			expectedError: err.IndexOutOfRange(2, 12),
		},
		{
			name:          "index less than 0",
			slice:         []int{123, 100},
			index:         -1,
			expectedError: err.IndexOutOfRange(2, -1),
		},
		{
			name:          "index last",
			slice:         []int{123, 100, 101, 102, 102, 102},
			index:         5,
			expectedSlice: []int{123, 100, 101, 102, 102},
			expectedVal:   102,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res, val, err := Delete(tc.slice, tc.index)
			assert.Equal(t, tc.expectedError, err)
			if err != nil {
				return
			}
			assert.Equal(t, tc.expectedSlice, res)
			assert.Equal(t, tc.expectedVal, val)
		})
	}
}
