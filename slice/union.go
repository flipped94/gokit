package slice

import "github.com/flipped94/gokit/set"

// Union 并集，只支持 comparable
func Union[T comparable](source []T, elements []T) []T {
	sourceSet := set.OfSet[T](source)
	for _, v := range elements {
		sourceSet.Add(v)
	}
	return sourceSet.Keys()
}

// UnionFunc 并集，自定义相等
func UnionFunc[T any](source []T, elements []T, equal equalFunc[T]) []T {
	var result = make([]T, 0, len(source)+len(elements))
	for _, v := range source {
		if !Contains(result, v, equal) {
			result = append(result, v)
		}
	}
	for _, v := range elements {
		if !Contains(result, v, equal) {
			result = append(result, v)
		}
	}
	return result
}
