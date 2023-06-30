package ds

type Stack[T any] struct {
	vals []T
}

func (s *Stack[T]) Push(val T) {
	s.vals = append(s.vals, val)
}

func (s *Stack[T]) Pop() (T, bool) {
	if len(s.vals) == 0 {

		// Tのゼロ値を帰す
		var zero T
		return zero, false
	}

	val := s.vals[len(s.vals)-1]
	s.vals = s.vals[:len(s.vals)-1]
	return val, true
}

// Size スタックのサイズを返す
func (s *Stack[T]) Size() int {
	return len(s.vals)
}

// IsEmpty スタックが空かどうかチェックする
func (s *Stack[T]) IsEmpty() bool {
	if len(s.vals) == 0 {
		return true
	} else {
		return false
	}
}
