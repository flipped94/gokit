package slice

// Map 转化
func Map[S any, T any](source []S, ap apply[S, T]) []T {
	results := make([]T, len(source))
	for i, s := range source {
		results[i] = ap(s)
	}
	return results
}
