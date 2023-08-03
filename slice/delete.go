package slice

import "github.com/flipped94/gokit/internal/slice"

// Delete 删除 index 处的元素
func Delete[T any](source []T, index int) ([]T, error) {
	res, _, err := slice.Delete[T](source, index)
	return res, err
}

// Filter 删除符合条件的元素
func Filter[T any](source []T, prd predicate[T]) []T {
	pos := 0
	for index := range source {
		// 判断是否满足删除的条件
		if prd(source[index]) {
			continue
		}
		// 移动元素
		source[pos] = source[index]
		pos++
	}
	return source[:pos]
}
