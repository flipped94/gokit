package slice

import (
	"math"

	"github.com/flipped94/gokit/set"
)

// Intersect 交集，只支持 comparable 类型
func Intersect[T comparable](source []T, elements []T) []T {
	ms := set.OfSet[T](source)
	counts := int(math.Min(float64(len(source)), float64(len(elements))))
	result := make([]T, 0, counts)
	// 交集小于等于两个集合中的任意一个
	for _, val := range elements {
		if ms.Exist(val) {
			result = append(result, val)
		}
	}
	return result
}

// Intersect 交集，自定义相等
func IntersectFunc[T any](source []T, elements []T, equal equalFunc[T]) []T {
	counts := int(math.Min(float64(len(source)), float64(len(elements))))
	result := make([]T, 0, counts)
	for _, valSrc := range source {
		for _, valDst := range elements {
			if equal(valDst, valSrc) {
				result = append(result, valSrc)
				break
			}
		}
	}
	return result
}
