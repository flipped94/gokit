package list

import (
	"github.com/flipped94/gokit/internal/err"
	"github.com/flipped94/gokit/internal/slice"
)

var (
	_ List[any] = &ArrayList[any]{}
)

// ArrayList 基于切片的简单封装
type ArrayList[T any] struct {
	elements []T
}

// NewArrayList 初始化一个 len 为 0，cap 为 cap 的ArrayList
func NewArrayList[T any](cap int) *ArrayList[T] {
	return &ArrayList[T]{elements: make([]T, 0, cap)}
}

// NewArrayListOf 直接使用 source，而不会执行复制
func NewArrayListOf[T any](source []T) *ArrayList[T] {
	return &ArrayList[T]{
		elements: source,
	}
}

// Get 返回对应下标的元素，
// 在下标超出范围的情况下，返回错误
func (list *ArrayList[T]) Get(index int) (T, error) {
	l := list.Len()
	if index < 0 || index >= l {
		var zero T
		return zero, err.IndexOutOfRange(l, index)
	}
	return list.elements[index], nil
}

// Append 往ArrayList里追加数据
func (list *ArrayList[T]) Append(elements ...T) error {
	list.elements = append(list.elements, elements...)
	return nil
}

// Add 在 ArrayList 下标为 index 的位置插入一个元素
// 当 index 等于 ArrayList 长度等同于 append
func (list *ArrayList[T]) Add(index int, element T) error {
	if index < 0 || index > len(list.elements) {
		return err.IndexOutOfRange(len(list.elements), index)
	}
	list.elements = append(list.elements, element)
	copy(list.elements[index+1:], list.elements[index:])
	list.elements[index] = element
	return nil
}

// Set 设置ArrayList里index位置的值为t
func (list *ArrayList[T]) Set(index int, element T) error {
	l := len(list.elements)
	if index >= l || index < 0 {
		return err.IndexOutOfRange(l, index)
	}
	list.elements[index] = element
	return nil
}

// Delete 方法会在必要的时候引起缩容，其缩容规则是：
// - 如果容量 > 2048，并且长度小于容量一半，那么就会缩容为原本的 5/8
// - 如果容量 (64, 2048]，如果长度是容量的 1/4，那么就会缩容为原本的一半
// - 如果此时容量 <= 64，那么我们将不会执行缩容。在容量很小的情况下，浪费的内存很少，所以没必要消耗 CPU去执行缩容
func (list *ArrayList[T]) Delete(index int) (T, error) {
	res, t, err := slice.Delete(list.elements, index)
	if err != nil {
		return t, err
	}
	list.elements = res
	list.shrink()
	return t, nil
}

// shrink 数组缩容
func (list *ArrayList[T]) shrink() {
	list.elements = slice.Shrink(list.elements)
}

// Len 返回长度
func (list *ArrayList[T]) Len() int {
	return len(list.elements)
}

// Cap 返回容量
func (list *ArrayList[T]) Cap() int {
	return cap(list.elements)
}

// Range 遍历 List 的所有元素
func (a *ArrayList[T]) Range(fn func(index int, t T) error) error {
	for key, value := range a.elements {
		e := fn(key, value)
		if e != nil {
			return e
		}
	}
	return nil
}

// ToSlice 将 List 转化为一个切片
// 不允许返回 nil，在没有元素的情况下，
// 必须返回一个长度和容量都为 0 的切片
// ToSlice 每次调用都必须返回一个全新的切片
func (list *ArrayList[T]) ToSlice() []T {
	res := make([]T, len(list.elements))
	copy(res, list.elements)
	return res
}
