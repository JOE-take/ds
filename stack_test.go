package ds

import (
	"testing"
)

func Test_stack(t *testing.T) {
	s := Stack[float64]{}

	if !s.IsEmpty() {
		t.Error("スタックの初期値が空でない")
	}

	s.Push(3.14)
	if result, _ := s.Pop(); result != 3.14 {
		t.Error("スタックから取り出した値が正常でない, 期待した結果: 3.14, 実際の値: ", result)
	}
	if !s.IsEmpty() {
		t.Error("実際にスタックが空の時 IsEmpty == true でない")
	}

	s.Push(0.1)
	s.Push(0.2)
	s.Push(0.3)
	if result, _ := s.Pop(); result != 0.3 {
		t.Error("スタックから取り出した値が正常でない, 期待した結果: 0.3, 実際の値: ", result)
	}
	if result := s.Size(); result != 2 {
		t.Error("Size() が適切な値を返していない, 期待した結果: 2, 実際の結果: ", result)
	}

	_, ok1 := s.Pop()
	_, ok2 := s.Pop()
	if ok1 && ok2 == false {
		t.Errorf("ok が正しく動いていない, ok1: %t, ok2: %t", ok1, ok2)
	}

}
