package set

type Set[T any] interface {
	// Add 添加元素
	Add(key T)
	// Delete 删除元素
	Delete(key T)
	// Exist 元素是否存在
	Exist(key T) bool
	// Keys 返回所有元素
	Keys() []T
}
