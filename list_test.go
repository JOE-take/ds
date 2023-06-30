package ds

import (
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
}
