package slice

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUnionSet(t *testing.T) {
	tests := []struct {
		name     string
		source   []int
		elements []int
		expect   []int
	}{
		{
			name:     "normal union",
			source:   []int{1, 1, 2, 3},
			elements: []int{4, 5, 6, 6},
			expect:   []int{1, 2, 3, 4, 5, 6},
		},
		{
			name:     "empty source",
			source:   []int{},
			elements: []int{1, 3},
			expect:   []int{1, 3},
		},
		{
			name:     "empty elements",
			source:   []int{1, 3},
			elements: []int{},
			expect:   []int{1, 3},
		},
		{
			name:     "both empty",
			source:   []int{},
			elements: []int{},
			expect:   []int{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := Union[int](tt.source, tt.elements)
			assert.ElementsMatch(t, tt.expect, res)
		})
	}
}

func TestUnionFunc(t *testing.T) {
	tests := []struct {
		name     string
		source   []int
		elements []int
		expect   []int
	}{
		{
			name:     "normal union",
			source:   []int{1, 1, 2, 3},
			elements: []int{4, 5, 6, 6},
			expect:   []int{1, 2, 3, 4, 5, 6},
		},
		{
			name:     "empty source",
			source:   []int{},
			elements: []int{1, 3},
			expect:   []int{1, 3},
		},
		{
			name:     "empty elements",
			source:   []int{1, 3},
			elements: []int{},
			expect:   []int{1, 3},
		},
		{
			name:     "both empty",
			source:   []int{},
			elements: []int{},
			expect:   []int{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := UnionFunc[int](tt.source, tt.elements, func(src, dst int) bool {
				return src == dst
			})
			assert.ElementsMatch(t, tt.expect, res)
		})
	}
}
