package slice

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestContains(t *testing.T) {
	tests := []struct {
		name   string
		source []int
		e      int
		expect bool
	}{
		{
			name:   "contains",
			source: []int{1, 2, 3, 4, 5},
			e:      4,
			expect: true,
		},
		{
			name:   "not contains",
			source: []int{1, 2, 3, 4, 5},
			e:      6,
			expect: false,
		},
		{
			name:   "empty source is 0",
			source: []int{},
			e:      1,
			expect: false,
		},
		{name: "nil source",
			e:      1,
			expect: false,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.expect, Contains[int](test.source, test.e, func(src, dst int) bool {
				return src == dst
			}))
		})
	}
}

func TestContainsAny(t *testing.T) {
	tests := []struct {
		name     string
		source   []int
		elements []int
		expect   bool
	}{
		{
			name:     "contains",
			source:   []int{1, 2, 3, 4, 5},
			elements: []int{1, 6},
			expect:   true,
		},
		{
			name:     "not contains",
			source:   []int{1, 2, 3, 4, 5},
			elements: []int{0, 6},
			expect:   false,
		},
		{
			name:     "empty source",
			source:   []int{},
			elements: []int{1},
			expect:   false,
		},
		{
			name:     "nil source",
			elements: []int{1},
			expect:   false,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.expect, ContainsAny[int](test.source, test.elements))
		})
	}
}

func TestContainsAll(t *testing.T) {
	tests := []struct {
		name     string
		source   []int
		elements []int
		expect   bool
	}{
		{
			name:     "contains all",
			source:   []int{1, 2, 3, 4, 5},
			elements: []int{1, 2, 3},
			expect:   true,
		},
		{
			name:     "contains",
			source:   []int{1, 2, 3, 4, 5},
			elements: []int{1, 2, 3, 6},
			expect:   false,
		},
		{
			name:     "empty source",
			source:   []int{},
			elements: []int{1},
			expect:   false,
		},
		{
			name:     "nil source and empty elements",
			source:   nil,
			elements: []int{},
			expect:   true,
		},
		{
			name:   "source and elements both nil",
			source: nil,
			expect: true,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.expect, ContainsAll[int](test.source, test.elements))
		})
	}
}
