package slice

import "github.com/flipped94/gokit/set"

// Contains 判断 source 是否包含 e
func Contains[T comparable](source []T, e T, equal equalFunc[T]) bool {
	for _, v := range source {
		if equal(v, e) {
			return true
		}
	}
	return false
}

// ContainsAny 判断 source 是否包含 elements 中的任何一个元素
func ContainsAny[T comparable](source []T, elements []T) bool {
	hs := set.OfSet[T](source)
	for _, v := range elements {
		if hs.Exist(v) {
			return true
		}
	}
	return false
}

// ContainsAll 判断 source 是否包含 elements 中的所有元素
func ContainsAll[T comparable](source []T, elements []T) bool {
	hs := set.OfSet[T](source)
	for _, v := range elements {
		if !hs.Exist(v) {
			return false
		}
	}
	return true
}
