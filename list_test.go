package ds

import (
	"reflect"
	"testing"
)

func Test_List_Integer(t *testing.T) {
	var l List[int]
	if !l.IsEmpty() {
		t.Error("リストは最初は空であるべきだが、そうでない")
	}

	if l.Size() != 0 {
		t.Error("リストの最初のサイズは0であるべきだが、そうでない")
	}

	l.Append(1)
	l.Append(2)
	l.Append(3)

	cur := l.Head()
	if cur.Get() != 1 {
		t.Errorf("Expected: 1, Actual: %d", cur.Get())
	}

	cur, _ = cur.Next()
	if cur.Get() != 2 {
		t.Errorf("Expected: 2, Actual: %d", cur.Get())
	}

	cur, _ = cur.Next()
	if cur.Get() != 3 {
		t.Errorf("Expected: 3, Actual: %d", cur.Get())
	}

	if _, b := cur.Next(); b {
		t.Errorf("Expected: false, Actual: %t", b)
	}

	cur = l.Head()
	l.Insert(cur, 0)
	if cur.Get() != 1 {
		t.Errorf("Expected: 1, Actual: %d", cur.Get())
	}

	cur, _ = cur.Next()
	if cur.Get() != 0 {
		t.Errorf("Expected: 0, Actual: %d", cur.Get())
	}

	cur, _ = cur.Next()
	if cur.Get() != 2 {
		t.Errorf("Expected: 2, Actual: %d", cur.Get())
	}

	l.Initialize()
	l.Append(1)
	l.Append(2)
	l.Append(3)
	l.Append(4)

	if l.Size() != 4 {
		t.Errorf("Expected: 4, Actual: %d", l.Size())
	}

	l.Remove(l.Head())
	if l.Size() != 3 {
		t.Errorf("Expected: 3, Actual: %d", l.Size())
	}

	test := []int{2, 3, 4}
	arr := l.ToSlice()
	if !reflect.DeepEqual(test, arr) {
		t.Errorf("Expected: %v, Actual: %v", test, arr)
	}

	cur = l.Head()
	cur, _ = cur.Next()
	l.Insert(cur, 0)
	test = []int{2, 3, 0, 4}
	arr = l.ToSlice()
	if !reflect.DeepEqual(test, arr) {
		t.Errorf("Expected: %v, Actual: %v", test, arr)
	}

}
