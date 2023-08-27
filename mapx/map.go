package mapx

// Keys 返回所有 key。
func Keys[K comparable, V any](m map[K]V) []K {
	res := make([]K, 0, len(m))
	for k := range m {
		res = append(res, k)
	}
	return res
}

// Values 返回所有 value。
func Values[K comparable, V any](m map[K]V) []V {
	res := make([]V, 0, len(m))
	for k := range m {
		res = append(res, m[k])
	}
	return res
}

// Entry 返回所有 key,value。
func Entries[K comparable, V any](m map[K]V) []Entry[K, V] {
	res := make([]Entry[K, V], 0, len(m))
	for k := range m {
		res = append(res, Entry[K, V]{
			Key:   k,
			Value: m[k],
		})

	}
	return res
}
