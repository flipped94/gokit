package mapx

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestKeys(t *testing.T) {
	testCases := []struct {
		name   string
		kvs    map[string]int
		expect []string
	}{
		{
			name:   "nil map",
			kvs:    nil,
			expect: []string{},
		},
		{
			name:   "empty map",
			kvs:    map[string]int{},
			expect: []string{},
		},
		{
			name: "map with key value",
			kvs: map[string]int{
				"1": 1,
				"2": 2,
				"3": 3,
			},
			expect: []string{"1", "2", "3"},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res := Keys[string, int](tc.kvs)
			assert.ElementsMatch(t, tc.expect, res)
		})
	}
}

func TestValues(t *testing.T) {
	testCases := []struct {
		name   string
		kvs    map[string]int
		expect []int
	}{
		{
			name:   "nil map",
			kvs:    nil,
			expect: []int{},
		},
		{
			name:   "empty map",
			kvs:    map[string]int{},
			expect: []int{},
		},
		{
			name: "map with key value",
			kvs: map[string]int{
				"1": 1,
				"2": 2,
				"3": 3,
			},
			expect: []int{1, 2, 3},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res := Values[string, int](tc.kvs)
			assert.ElementsMatch(t, tc.expect, res)
		})
	}
}

func TestEntries(t *testing.T) {
	testCases := []struct {
		name   string
		kvs    map[string]int
		expect []Entry[string, int]
	}{
		{
			name:   "nil",
			kvs:    nil,
			expect: []Entry[string, int]{},
		},
		{
			name:   "empty",
			kvs:    map[string]int{},
			expect: []Entry[string, int]{},
		},
		{
			name: "map with key value",
			kvs: map[string]int{
				"1": 1,
				"2": 2,
				"3": 3,
			},
			expect: []Entry[string, int]{
				{
					Key:   "1",
					Value: 1,
				},
				{
					Key:   "2",
					Value: 2,
				},
				{
					Key:   "3",
					Value: 3,
				},
			},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			keys := Entries[string, int](tc.kvs)
			assert.ElementsMatch(t, tc.expect, keys)
		})
	}
}
