package queue

import (
	"errors"

	"github.com/flipped94/gokit"
	"github.com/flipped94/gokit/internal/slice"
)

var ErrEmptyQueue = errors.New("gokit: 队列为空")

var ErrOutOfCapacity = errors.New("gokit: 超出最大容量限制")

// PriorityQueue 是一个基于小顶堆的优先队列
// 当 capacity <= 0 时，为无界队列，容量会动态扩缩容
// 当 capacity > 0 时，为有界队列，初始化后不会扩缩容
type PriorityQueue[T any] struct {

	// 比较元素
	compare gokit.Comparator[T]

	// 队列容量
	capacity int

	// 队列元素
	data []T
}

// NewPriorityQueue 创建优先队列
func NewPriorityQueue[T any](capacity int, comparator gokit.Comparator[T]) *PriorityQueue[T] {
	sliceCap := capacity + 1
	if capacity < 1 {
		capacity = 0
		sliceCap = 64
	}
	return &PriorityQueue[T]{
		capacity: capacity,
		data:     make([]T, 1, sliceCap),
		compare:  comparator,
	}
}

// Len 队列长度
func (p *PriorityQueue[T]) Len() int {
	return len(p.data) - 1
}

// Cap 无界队列返回0，有界队列返回创建队列时设置的值
func (p *PriorityQueue[T]) Cap() int {
	return p.capacity
}

// Peek 查看堆根节点
func (p *PriorityQueue[T]) Peek() (T, error) {
	if p.isEmpty() {
		var t T
		return t, ErrEmptyQueue
	}
	return p.data[1], nil
}

// Enqueue 元素入队
func (p *PriorityQueue[T]) Enqueue(t T) error {
	if p.isFull() {
		return ErrOutOfCapacity
	}

	p.data = append(p.data, t)
	node, parent := len(p.data)-1, (len(p.data)-1)/2
	for parent > 0 && p.compare(p.data[node], p.data[parent]) < 0 {
		p.data[parent], p.data[node] = p.data[node], p.data[parent]
		node = parent
		parent = parent / 2
	}

	return nil
}

// Dequeue 元素出队
func (p *PriorityQueue[T]) Dequeue() (T, error) {
	if p.isEmpty() {
		var t T
		return t, ErrEmptyQueue
	}

	pop := p.data[1]
	p.data[1] = p.data[len(p.data)-1]
	p.data = p.data[:len(p.data)-1]
	p.shrinkIfNecessary()
	p.heapify(p.data, len(p.data)-1, 1)
	return pop, nil
}

// isFull 判断队列是否已满
func (p *PriorityQueue[T]) isFull() bool {
	return p.capacity > 0 && len(p.data)-1 == p.capacity
}

// isEmpty 判断队列是否为空
func (p *PriorityQueue[T]) isEmpty() bool {
	return len(p.data) < 2
}

// IsBoundless 是否有界
func (p *PriorityQueue[T]) IsBoundless() bool {
	return p.capacity <= 0
}

// shrinkIfNecessary 缩容
func (p *PriorityQueue[T]) shrinkIfNecessary() {
	if p.IsBoundless() {
		p.data = slice.Shrink[T](p.data)
	}
}

// heapify 调整堆结构
func (p *PriorityQueue[T]) heapify(data []T, n, i int) {
	minPos := i
	for {
		if left := i * 2; left <= n && p.compare(data[left], data[minPos]) < 0 {
			minPos = left
		}
		if right := i*2 + 1; right <= n && p.compare(data[right], data[minPos]) < 0 {
			minPos = right
		}
		if minPos == i {
			break
		}
		data[i], data[minPos] = data[minPos], data[i]
		i = minPos
	}
}
