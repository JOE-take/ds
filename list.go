package ds

import (
	"fmt"
)

type Node[T any] struct {
	value T
	next  *Node[T]
	prev  *Node[T]
}

// List 双方向連結リスト; Doubly Linked List
type List[T any] struct {
	size int
	head *Node[T]
	tail *Node[T]
}

// Get 現在要素を返す
func (node *Node[T]) Get() T {
	return node.value
}

// Next 現在ノードの次のノードを返す。
// 現在ノードが末尾のノードの時のみ false を返す。
func (node *Node[T]) Next() (*Node[T], bool) {
	if node.next == nil {
		return nil, false
	} else {
		return node.next, true
	}
}

// Prev 現在ノードの前のノードを返す。
// 現在ノードが最初のノードの時のみ false を返す。
func (node *Node[T]) Prev() (*Node[T], bool) {
	if node.prev == nil {
		return nil, false
	} else {
		return node.prev, true
	}
}

// Size リストのサイズを返す。
func (l *List[T]) Size() int {
	return l.size
}

// Head リストの最初の要素を返す。
func (l *List[T]) Head() *Node[T] {
	return l.head
}

// Tail リストの末尾の要素を返す。
func (l *List[T]) Tail() *Node[T] {
	return l.tail
}

// IsEmpty リストが空かどうかを返す。
func (l *List[T]) IsEmpty() bool {
	return l.size == 0
}

func (l *List[T]) Initialize() {
	l.head = nil
	l.tail = nil
	l.size = 0
}

// Append リストの末尾に新しいノードを追加し、それを末尾とする。
func (l *List[T]) Append(element T) {
	newNode := Node[T]{
		value: element,
		next:  nil,
		prev:  l.tail,
	}
	if l.IsEmpty() {
		l.head = &newNode
		l.tail = &newNode
		l.size++
	} else {
		l.tail.next = &newNode
		l.tail = &newNode
		l.size++
	}
}

// Insert 現在ノードの次に新しいノードを挿入
func (l *List[T]) Insert(node *Node[T], element T) {
	if node.next == nil {
		l.Append(element)
		return
	}
	newNode := Node[T]{
		value: element,
		next:  node.next,
		prev:  node,
	}
	node.next.prev = &newNode
	node.next = &newNode
	l.size++
}

// Remove リストからノードを削除
func (l *List[T]) Remove(node *Node[T]) {
	switch {
	case node.next == nil && node.prev == nil:

	case node.next == nil:
		node.prev.next = nil
		l.tail = node.prev
	case node.prev == nil:
		node.next.prev = nil
		l.head = node.next
	default:
		node.next.prev = node.prev
		node.prev.next = node.next
	}
	l.size--
}

// ShowAll リストの全ての要素をスペース区切りで表示
func (l *List[T]) ShowAll() {
	if l.size == 0 {
		return
	}
	cur := l.head
	for cur != nil {
		fmt.Print(cur.value, " ")
		cur = cur.next
	}
	fmt.Println()
}

// ToSlice リストをスライスに変換
func (l *List[T]) ToSlice() []T {
	arr := make([]T, l.size)

	cur := l.head
	for i := 0; cur != nil; i++ {
		arr[i] = cur.value
		cur = cur.next
	}

	return arr
}
