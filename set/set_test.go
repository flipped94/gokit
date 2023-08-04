package set

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSet_Add(t *testing.T) {
	added := []int{1, 2, 3, 1}
	s := NewMapSet[int](10)
	t.Run("Add", func(t *testing.T) {
		for _, val := range added {
			s.Add(val)
		}
		assert.Equal(t, s.m, map[int]struct{}{
			1: {},
			2: {},
			3: {},
		})
	})
}

func TestSet_Delete(t *testing.T) {
	testcases := []struct {
		name    string
		deleted int
		source  map[int]struct{}
		expect  map[int]struct{}
	}{
		{
			name:    "Delete",
			deleted: 2,
			source: map[int]struct{}{
				2: {},
			},
			expect: map[int]struct{}{},
		},
		{
			name:    "deleted not found",
			deleted: 3,
			source: map[int]struct{}{
				2: {},
			},
			expect: map[int]struct{}{
				2: {},
			},
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			s := NewMapSet[int](10)
			s.m = tc.source
			s.Delete(tc.deleted)
			assert.Equal(t, tc.expect, s.m)
		})
	}
}

func TestSet_Exist(t *testing.T) {
	s := NewMapSet[int](10)
	s.Add(1)
	testcases := []struct {
		name  string
		val   int
		exist bool
	}{
		{
			name:  "found",
			val:   1,
			exist: true,
		},
		{
			name:  "not fonud",
			val:   2,
			exist: false,
		},
	}
	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			ok := s.Exist(tc.val)
			assert.Equal(t, tc.exist, ok)
		})
	}
}

func TestSet_Values(t *testing.T) {
	s := NewMapSet[int](10)
	testcases := []struct {
		name   string
		source map[int]struct{}
		expect map[int]struct{}
	}{
		{
			name: "found values",
			source: map[int]struct{}{
				1: {},
				2: {},
				3: {},
			},
			expect: map[int]struct{}{
				1: {},
				2: {},
				3: {},
			},
		},
	}
	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			s.m = tc.source
			vals := s.Keys()
			ok := equal(vals, tc.expect)
			assert.Equal(t, true, ok)
		})
	}
}

func equal(nums []int, m map[int]struct{}) bool {
	for _, num := range nums {
		_, ok := m[num]
		if !ok {
			return false
		}
		delete(m, num)
	}
	return true && len(m) == 0
}
