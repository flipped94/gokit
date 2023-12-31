package slice

// equalFunc 比较两个元素是否相等
type equalFunc[T any] func(src, dst T) bool

// 判断 e 是否满足条件
type predicate[T any] func(e T) bool

// 类型 S 到类型 T 的映射
type apply[S any, T any] func(soure S) T
