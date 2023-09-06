package queue

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/flipped94/gokit"
)

func TestNewPriorityQueue(t *testing.T) {
	data := []int{6, 5, 4, 3, 2, 1}
	testCases := []struct {
		name     string
		queue    *PriorityQueue[int]
		capacity int
		data     []int
		expected []int
	}{
		{
			name:     "无界队列",
			queue:    NewPriorityQueue(0, comparator()),
			capacity: 0,
			data:     data,
			expected: []int{1, 2, 3, 4, 5, 6},
		},
		{
			name:     "有界队列 ",
			queue:    NewPriorityQueue(len(data), comparator()),
			capacity: len(data),
			data:     data,
			expected: []int{1, 2, 3, 4, 5, 6},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, 0, tc.queue.Len())
			for _, d := range data {
				err := tc.queue.Enqueue(d)
				assert.NoError(t, err)
				if err != nil {
					return
				}
			}
			assert.Equal(t, tc.capacity, tc.queue.Cap())
			assert.Equal(t, len(data), tc.queue.Len())
			res := make([]int, 0, len(data))
			for tc.queue.Len() > 0 {
				el, err := tc.queue.Dequeue()
				assert.NoError(t, err)
				if err != nil {
					return
				}
				res = append(res, el)
			}
			assert.Equal(t, tc.expected, res)
		})

	}

}

func TestPriorityQueue_Peek(t *testing.T) {
	testCases := []struct {
		name     string
		capacity int
		data     []int
		wantErr  error
	}{
		{
			name:     "非空队列",
			capacity: 0,
			data:     []int{6, 5, 4, 3, 2, 1},
			wantErr:  ErrEmptyQueue,
		},
		{
			name:     "空队列",
			capacity: 0,
			data:     []int{},
			wantErr:  ErrEmptyQueue,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			queue := NewPriorityQueue[int](tc.capacity, comparator())
			for _, el := range tc.data {
				err := queue.Enqueue(el)
				require.NoError(t, err)
			}
			for queue.Len() > 0 {
				peek, err := queue.Peek()
				assert.NoError(t, err)
				el, _ := queue.Dequeue()
				assert.Equal(t, el, peek)
			}
			_, err := queue.Peek()
			assert.Equal(t, tc.wantErr, err)
		})

	}
}

func TestPriorityQueue_Dequeue(t *testing.T) {
	testCases := []struct {
		name      string
		data      []int
		wantErr   error
		wantVal   int
		wantSlice []int
	}{
		{
			name:    "空队列",
			data:    []int{},
			wantErr: ErrEmptyQueue,
		},
		{
			name:      "只有一个元素",
			data:      []int{10},
			wantVal:   10,
			wantSlice: []int{0},
		},
		{
			name:      "many",
			data:      []int{6, 5, 4, 3, 2, 1},
			wantVal:   1,
			wantSlice: []int{0, 2, 3, 5, 6, 4},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			q := priorityQueueOf(0, tc.data, comparator())
			require.NotNil(t, q)
			val, err := q.Dequeue()
			assert.Equal(t, tc.wantErr, err)
			if err != nil {
				return
			}
			assert.Equal(t, tc.wantSlice, q.data)
			assert.Equal(t, tc.wantVal, val)
		})
	}
}

func priorityQueueOf(capacity int, data []int, compare gokit.Comparator[int]) *PriorityQueue[int] {
	q := NewPriorityQueue[int](capacity, compare)
	for _, el := range data {
		err := q.Enqueue(el)
		if err != nil {
			return nil
		}
	}
	return q
}

func comparator() gokit.Comparator[int] {
	return func(a, b int) int {
		if a < b {
			return -1
		}
		if a == b {
			return 0
		}
		return 1
	}
}
