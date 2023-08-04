package slice

import "github.com/flipped94/gokit/internal/err"

func Add[T any](source []T, e T, index int) ([]T, error) {
	length := len(source)
	// 越界检查
	if source == nil || index < 0 || index >= length {
		return nil, err.IndexOutOfRange(length, index)
	}
	// 先追加一个零值
	var zeroValue T
	source = append(source, zeroValue)
	// 向后移动
	for i := len(source) - 1; i > index; i-- {
		if i-1 >= 0 {
			source[i] = source[i-1]
		}
	}
	// 赋值
	source[index] = e
	return source, nil
}
