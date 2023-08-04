package set

type MapSet[T comparable] struct {
	m map[T]struct{}
}

func NewMapSet[T comparable](size int) *MapSet[T] {
	return &MapSet[T]{
		m: make(map[T]struct{}, size),
	}
}

func OfSet[T comparable](source []T) *MapSet[T] {
	temp := make(map[T]struct{}, len(source))
	for _, v := range source {
		temp[v] = struct{}{}
	}
	return &MapSet[T]{
		m: temp,
	}
}

func (set *MapSet[T]) Add(key T) {
	set.m[key] = struct{}{}
}

func (set *MapSet[T]) Delete(key T) {
	delete(set.m, key)
}

func (set *MapSet[T]) Exist(key T) bool {
	_, ok := set.m[key]
	return ok
}

func (set *MapSet[T]) Keys() []T {
	keys := make([]T, 0, len(set.m))
	for key := range set.m {
		keys = append(keys, key)
	}
	return keys
}
