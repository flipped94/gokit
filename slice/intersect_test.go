package slice

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIntersect(t *testing.T) {
	tests := []struct {
		name     string
		source   []int
		elements []int
		expect   []int
	}{
		{
			name:     "normal intersect",
			source:   []int{1, 2, 3, 4},
			elements: []int{1, 3, 5},
			expect:   []int{1, 3},
		},
		{
			name:     "nil source",
			source:   nil,
			elements: []int{1, 3, 5},
			expect:   []int{},
		},
		{
			name:     "empty source",
			source:   []int{1, 2, 3},
			elements: []int{},
			expect:   []int{},
		},
		{
			name:     "empty elements",
			source:   []int{1, 3, 5},
			elements: []int{},
			expect:   []int{},
		},
		{
			name:     "nil elements",
			source:   []int{1, 3, 5, 5},
			elements: nil,
			expect:   []int{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := Intersect[int](tt.source, tt.elements)
			assert.ElementsMatch(t, tt.expect, res)
		})
	}
}

func TestIntersectFunc(t *testing.T) {
	tests := []struct {
		name     string
		source   []int
		elements []int
		expect   []int
	}{
		{
			name:     "normal intersect",
			source:   []int{1, 2, 3, 4},
			elements: []int{1, 3, 5},
			expect:   []int{1, 3},
		},
		{
			name:     "nil source",
			source:   nil,
			elements: []int{1, 3, 5},
			expect:   []int{},
		},
		{
			name:     "empty source",
			source:   []int{1, 2, 3},
			elements: []int{},
			expect:   []int{},
		},
		{
			name:     "empty elements",
			source:   []int{1, 3, 5},
			elements: []int{},
			expect:   []int{},
		},
		{
			name:     "nil elements",
			source:   []int{1, 3, 5, 5},
			elements: nil,
			expect:   []int{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := IntersectFunc[int](tt.source, tt.elements, func(src, dst int) bool {
				return src == dst
			})
			assert.ElementsMatch(t, tt.expect, res)
		})
	}
}
