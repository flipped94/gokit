package slice

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMap(t *testing.T) {
	tests := []struct {
		name   string
		source []int
		want   []string
	}{
		{
			name: "nil source",
			want: []string{},
		},
		{
			name:   "empty source",
			source: []int{},
			want:   []string{},
		},
		{
			name:   "int to string",
			source: []int{1, 2, 3},
			want:   []string{"1", "2", "3"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := Map(tt.source, func(src int) string {
				return strconv.Itoa(src)
			})
			assert.Equal(t, res, tt.want)
		})
	}
}
