package slice

// Index 返回和 e 相等的第一个元素下标
// -1 表示没找到
func Index[T comparable](source []T, e T) int {
	return IndexFunc[T](source, e, func(source, e T) bool {
		return source == e
	})
}

// IndexFunc 返回和 e 相等的第一个元素下标
// -1 表示没找到
func IndexFunc[T any](source []T, e T, equal equalFunc[T]) int {
	for k, v := range source {
		if equal(v, e) {
			return k
		}
	}
	return -1
}

// LastIndex 返回和 e 相等的最后一个元素下标
// -1 表示没找到
func LastIndex[T comparable](source []T, e T) int {
	return LastIndexFunc[T](source, e, func(source, e T) bool {
		return source == e
	})
}

// LastIndexFunc 返回和 e 相等的最后一个元素下标
// -1 表示没找到
func LastIndexFunc[T any](source []T, e T, equal equalFunc[T]) int {
	for i := len(source) - 1; i >= 0; i-- {
		if equal(e, source[i]) {
			return i
		}
	}
	return -1
}

// IndexAll 返回和 e 相等的所有元素的下标
func IndexAll[T comparable](source []T, e T) []int {
	return IndexAllFunc[T](source, e, func(source, e T) bool {
		return source == e
	})
}

// IndexAllFunc 返回和 e 相等的所有元素的下标
func IndexAllFunc[T any](source []T, e T, equal equalFunc[T]) []int {
	var indexes = make([]int, 0, len(source))
	for k, v := range source {
		if equal(v, e) {
			indexes = append(indexes, k)
		}
	}
	return indexes
}
