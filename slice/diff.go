package slice

import "github.com/flipped94/gokit/set"

// Diff 差集，只支持 comparable 类型
func Diff[T comparable](source []T, elements []T) []T {
	ms := set.OfSet[T](source)
	for _, val := range elements {
		ms.Delete(val)
	}
	return ms.Keys()
}

// DiffFunc 差集，自定义相等
func DiffFunc[T any](source []T, elements []T, equal equalFunc[T]) []T {
	results := make([]T, 0, len(source))
	for _, v := range source {
		if !Contains[T](elements, v, equal) {
			results = append(results, v)
		}
	}
	return results
}
