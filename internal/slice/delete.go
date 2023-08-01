package slice

import "github.com/flipped94/gokit/internal/err"

func Delete[T any](src []T, index int) ([]T, T, error) {
	var res T
	length := len(src)
	if src == nil || index < 0 || index >= length {
		return nil, res, err.IndexOutOfRange(length, index)
	}
	res = src[index]
	for i := index; i < length-1; i++ {
		src[i] = src[i+1]
	}
	src = src[:length-1]
	return src, res, nil
}
