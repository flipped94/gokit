package slice

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/flipped94/gokit/internal/err"
)

func TestDelete(t *testing.T) {
	testCases := []struct {
		name    string
		source  []int
		index   int
		result  []int
		wantErr error
	}{
		{
			name:   "index 0",
			source: []int{1, 2, 3},
			index:  0,
			result: []int{2, 3},
		},
		{
			name:    "index -1",
			source:  []int{1, 2, 3},
			index:   -1,
			wantErr: err.IndexOutOfRange(3, -1),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res, err := Delete(tc.source, tc.index)
			assert.Equal(t, tc.wantErr, err)
			if err != nil {
				return
			}
			assert.Equal(t, tc.result, res)
		})
	}
}

func TestFilter(t *testing.T) {

	testCases := []struct {
		name   string
		source []int
		prd    predicate[int]
		result []int
	}{
		{
			name:   "空切片",
			source: []int{},
			prd: func(e int) bool {
				return false
			},

			result: []int{},
		},
		{
			name:   "不删除元素",
			source: []int{0, 1, 2, 3, 4, 5, 6, 7},
			prd: func(e int) bool {
				return false
			},

			result: []int{0, 1, 2, 3, 4, 5, 6, 7},
		},
		{
			name:   "删除单个元素",
			source: []int{0, 1, 2, 3, 4, 5, 6, 7},
			prd: func(e int) bool {
				return e == 3
			},

			result: []int{0, 1, 2, 4, 5, 6, 7},
		},
		{
			name:   "删除多个不连续元素",
			source: []int{0, 1, 2, 3, 4, 5, 6, 7},
			prd: func(e int) bool {
				return e == 2 || e == 4
			},

			result: []int{0, 1, 3, 5, 6, 7},
		},
		{
			name:   "删除多个连续元素",
			source: []int{0, 1, 2, 3, 4, 5, 6, 7},
			prd: func(e int) bool {
				return e == 3 || e == 4
			},

			result: []int{0, 1, 2, 5, 6, 7},
		},
		{
			name:   "删除末尾元素",
			source: []int{0, 1, 2, 3, 4, 5, 6, 7},
			prd: func(e int) bool {
				return e == 7
			},

			result: []int{0, 1, 2, 3, 4, 5, 6},
		},
		{
			name:   "删除所有元素",
			source: []int{0, 1, 2, 3, 4, 5, 6, 7},
			prd: func(e int) bool {
				return true
			},

			result: []int{},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res := Filter(tc.source, tc.prd)
			assert.Equal(t, tc.result, res)
		})
	}
}
