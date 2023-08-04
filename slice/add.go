package slice

import "github.com/flipped94/gokit/internal/slice"

// Add 在 index 处添加元素
func Add[T any](source []T, e T, index int) ([]T, error) {
	res, err := slice.Add[T](source, e, index)
	return res, err
}
