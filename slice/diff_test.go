package slice

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDiff(t *testing.T) {
	tests := []struct {
		name     string
		source   []int
		elements []int
		expect   []int
	}{
		{
			name:     "normal diff",
			source:   []int{1, 2, 3, 4},
			elements: []int{1, 2, 3},
			expect:   []int{4},
		},
		{
			name:     "nil source",
			source:   nil,
			elements: []int{1, 2, 3},
			expect:   []int{},
		},
		{
			name:     "empty source",
			source:   []int{},
			elements: []int{1, 2, 3},
			expect:   []int{},
		},
		{
			name:     "nil elements",
			source:   []int{1, 2, 3, 4, 5},
			elements: nil,
			expect:   []int{1, 2, 3, 4, 5},
		},
		{
			name:     "empty elements",
			source:   []int{1, 2, 3, 4, 5},
			elements: []int{},
			expect:   []int{1, 2, 3, 4, 5},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := Diff[int](tt.source, tt.elements)
			assert.ElementsMatch(t, tt.expect, res)
		})
	}
}

func TestDiffFunc(t *testing.T) {
	tests := []struct {
		name     string
		source   []int
		elements []int
		expect   []int
	}{
		{
			name:     "normal diff",
			source:   []int{1, 2, 3, 4},
			elements: []int{1, 2, 3},
			expect:   []int{4},
		},
		{
			name:     "nil source",
			source:   nil,
			elements: []int{1, 2, 3},
			expect:   []int{},
		},
		{
			name:     "empty source",
			source:   []int{},
			elements: []int{1, 2, 3},
			expect:   []int{},
		},
		{
			name:     "nil elements",
			source:   []int{1, 2, 3, 4, 5},
			elements: nil,
			expect:   []int{1, 2, 3, 4, 5},
		},
		{
			name:     "empty elements",
			source:   []int{1, 2, 3, 4, 5},
			elements: []int{},
			expect:   []int{1, 2, 3, 4, 5},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := DiffFunc[int](tt.source, tt.elements, func(src, dst int) bool {
				return src == dst
			})
			assert.ElementsMatch(t, tt.expect, res)
		})
	}
}
