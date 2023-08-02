package slice

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIndex(t *testing.T) {
	testCases := []struct {
		name   string
		source []int
		e      int
		want   int
	}{
		{
			name:   "first one",
			source: []int{1, 1, 2, 3},
			e:      1,
			want:   0,
		},
		{
			name:   "empty source",
			source: []int{},
			e:      1,
			want:   -1,
		},
		{
			name: "nil source",
			e:    1,
			want: -1,
		},
		{
			name:   "e not exist",
			source: []int{1, 2, 3},
			e:      7,
			want:   -1,
		},
		{
			name:   "last one",
			source: []int{1, 2, 3, 4, 0},
			e:      0,
			want:   4,
		},
	}
	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.want, Index[int](test.source, test.e))
		})
	}
}

func TestIndexFunc(t *testing.T) {
	testCases := []struct {
		name   string
		source []int
		e      int
		want   int
	}{
		{
			name:   "first one",
			source: []int{1, 1, 2, 3},
			e:      1,
			want:   0,
		},
		{
			name:   "empty source",
			source: []int{},
			e:      1,
			want:   -1,
		},
		{
			name: "nil source",
			e:    1,
			want: -1,
		},
		{
			name:   "e not exist",
			source: []int{1, 2, 3},
			e:      7,
			want:   -1,
		},
		{
			name:   "last one",
			source: []int{1, 2, 3, 4, 0},
			e:      0,
			want:   4,
		},
	}
	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.want, IndexFunc[int](test.source, test.e, func(source, e int) bool {
				return source == e
			}))
		})
	}
}

func TestLastIndex(t *testing.T) {
	testCases := []struct {
		name   string
		source []int
		e      int
		want   int
	}{
		{
			name:   "last one 1",
			source: []int{1, 1, 2, 3},
			e:      1,
			want:   1,
		},
		{
			name:   "empty source",
			source: []int{},
			e:      1,
			want:   -1,
		},
		{
			name: "nil source",
			e:    1,
			want: -1,
		},
		{
			source: []int{1, 2, 3},
			e:      7,
			want:   -1,
			name:   "e not exist",
		},
		{
			source: []int{0, 1, 2, 3, 4, 0},
			e:      0,
			want:   5,
			name:   "last one 2",
		},
	}
	for _, test := range testCases {
		assert.Equal(t, test.want, LastIndex[int](test.source, test.e))
	}
}

func TestLastIndexFunc(t *testing.T) {
	testCases := []struct {
		name   string
		source []int
		e      int
		want   int
	}{
		{
			name:   "last one 1",
			source: []int{1, 1, 2, 3},
			e:      1,
			want:   1,
		},
		{
			name:   "empty source",
			source: []int{},
			e:      1,
			want:   -1,
		},
		{
			name: "nil source",
			e:    1,
			want: -1,
		},
		{
			source: []int{1, 2, 3},
			e:      7,
			want:   -1,
			name:   "e not exist",
		},
		{
			source: []int{0, 1, 2, 3, 4, 0},
			e:      0,
			want:   5,
			name:   "last one 2",
		},
	}
	for _, test := range testCases {
		assert.Equal(t, test.want, LastIndexFunc[int](test.source, test.e, func(source, e int) bool {
			return source == e
		}))
	}
}

func TestIndexAll(t *testing.T) {
	testCases := []struct {
		name   string
		source []int
		e      int
		want   []int
	}{
		{
			name:   "normal test",
			source: []int{1, 1, 3, 5},
			e:      1,
			want:   []int{0, 1},
		},
		{
			source: []int{},
			e:      1,
			want:   []int{},
			name:   "empty source",
		},
		{
			source: []int{1, 4, 6},
			e:      7,
			want:   []int{},
			name:   "e not exist",
		},
	}
	for _, test := range testCases {
		res := IndexAll[int](test.source, test.e)
		assert.ElementsMatch(t, test.want, res)
	}
}

func TestIndexAllFunc(t *testing.T) {
	testCases := []struct {
		name   string
		source []int
		e      int
		want   []int
	}{
		{
			name:   "normal test",
			source: []int{1, 1, 3, 5},
			e:      1,
			want:   []int{0, 1},
		},
		{
			source: []int{},
			e:      1,
			want:   []int{},
			name:   "empty source",
		},
		{
			source: []int{1, 4, 6},
			e:      7,
			want:   []int{},
			name:   "e not exist",
		},
	}
	for _, test := range testCases {
		res := IndexAllFunc[int](test.source, test.e, func(source, e int) bool {
			return source == e
		})
		assert.ElementsMatch(t, test.want, res)
	}
}
