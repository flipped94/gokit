package slice

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/flipped94/gokit/internal/err"
)

func TestAdd(t *testing.T) {
	testCases := []struct {
		name    string
		source  []int
		added   int
		index   int
		expect  []int
		wantErr error
	}{
		{
			name:   "index 0",
			source: []int{1, 2},
			added:  3,
			index:  0,
			expect: []int{3, 1, 2},
		},
		{
			name:    "index -1",
			source:  []int{1, 2},
			index:   -1,
			wantErr: err.IndexOutOfRange(2, -1),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res, err := Add(tc.source, tc.added, tc.index)
			assert.Equal(t, tc.wantErr, err)
			if err != nil {
				return
			}
			assert.Equal(t, tc.expect, res)
		})
	}
}
