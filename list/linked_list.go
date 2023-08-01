package list

import "github.com/flipped94/gokit/internal/err"

var (
	_ List[any] = &LinkedList[any]{}
)

// node 双向循环链表结点
type node[T any] struct {
	prev  *node[T]
	next  *node[T]
	value T
}

// LinkedList 双向循环链表
type LinkedList[T any] struct {
	head   *node[T]
	tail   *node[T]
	length int
}

// NewLinkedList 创建一个双向循环链表
func NewLinkedList[T any]() *LinkedList[T] {
	head := &node[T]{}
	tail := &node[T]{next: head, prev: head}
	head.next, head.prev = tail, tail
	return &LinkedList[T]{
		head: head,
		tail: tail,
	}
}

// NewLinkedListOf 将切片转换为双向循环链表, 直接使用了切片元素的值，而没有进行复制
func NewLinkedListOf[T any](source []T) *LinkedList[T] {
	list := NewLinkedList[T]()
	if err := list.Append(source...); err != nil {
		panic(err)
	}
	return list
}

// Get 返回对应下标的元素，
// 在下标超出范围的情况下，返回错误
func (list *LinkedList[T]) Get(index int) (T, error) {
	if !list.checkIndex(index) {
		var zeroValue T
		return zeroValue, err.IndexOutOfRange(list.Len(), index)
	}
	n := list.findNode(index)
	return n.value, nil
}

// Append 往链表最后添加元素
func (list *LinkedList[T]) Append(elements ...T) error {
	for _, t := range elements {
		node := &node[T]{prev: list.tail.prev, next: list.tail, value: t}
		node.prev.next, node.next.prev = node, node
		list.length++
	}
	return nil
}

// Add 在 LinkedList 下标为 index 的位置插入一个元素
// 当 index 等于 LinkedList 长度等同于 Append
func (list *LinkedList[T]) Add(index int, e T) error {
	if index < 0 || index > list.length {
		return err.IndexOutOfRange(list.length, index)
	}
	if index == list.length {
		return list.Append(e)
	}
	next := list.findNode(index)
	node := &node[T]{prev: next.prev, next: next, value: e}
	node.prev.next, node.next.prev = node, node
	list.length++
	return nil
}

// Set 设置链表中index索引处的值为 e
func (list *LinkedList[T]) Set(index int, e T) error {
	if !list.checkIndex(index) {
		return err.IndexOutOfRange(list.Len(), index)
	}
	node := list.findNode(index)
	node.value = e
	return nil
}

func (list *LinkedList[T]) findNode(index int) *node[T] {
	var cur *node[T]
	// 前半部分找
	if index <= list.Len()/2 {
		cur = list.head
		for i := -1; i < index; i++ {
			cur = cur.next
		}
	} else {
		// 前半部分找
		cur = list.tail
		for i := list.Len(); i > index; i-- {
			cur = cur.prev
		}
	}

	return cur
}

// Delete 删除指定位置的元素
func (list *LinkedList[T]) Delete(index int) (T, error) {
	if !list.checkIndex(index) {
		var zeroValue T
		return zeroValue, err.IndexOutOfRange(list.Len(), index)
	}
	node := list.findNode(index)
	node.prev.next = node.next
	node.next.prev = node.prev
	node.prev, node.next = nil, nil
	list.length--
	return node.value, nil
}

func (list *LinkedList[T]) checkIndex(index int) bool {
	return 0 <= index && index < list.Len()
}

// Len 返回长度
func (list *LinkedList[T]) Len() int {
	return list.length
}

// Cap 返回容量
func (list *LinkedList[T]) Cap() int {
	return list.Len()
}

func (list *LinkedList[T]) Range(fn func(index int, t T) error) error {
	for cur, i := list.head.next, 0; i < list.length; i++ {
		err := fn(i, cur.value)
		if err != nil {
			return err
		}
		cur = cur.next
	}
	return nil
}

func (list *LinkedList[T]) ToSlice() []T {
	slice := make([]T, list.length)
	for cur, i := list.head.next, 0; i < list.length; i++ {
		slice[i] = cur.value
		cur = cur.next
	}
	return slice
}
