package err

import "fmt"

func IndexOutOfRange(length int, index int) error {
	return fmt.Errorf("gokit: out of range: length: {%d}, and index: {%d}", length, index)
}
