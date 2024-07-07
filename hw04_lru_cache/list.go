package hw04lrucache

type List interface {
	Len() int
	Front() *ListItem
	Back() *ListItem
	PushFront(v interface{}) *ListItem
	PushBack(v interface{}) *ListItem
	Remove(i *ListItem)
	MoveToFront(i *ListItem)
}

type ListItem struct {
	Value interface{}
	Next  *ListItem
	Prev  *ListItem
}

type list struct {
	Length int
	Head   *ListItem
	Tail   *ListItem
}

func (l *list) Len() int {
	return l.Length
}

func (l *list) Front() *ListItem {
	return l.Tail
}

func (l *list) Back() *ListItem {
	return l.Head
}

func (l *list) PushBack(val interface{}) *ListItem {
	item := &ListItem{Prev: l.Head, Value: val, Next: nil}
	if l.Head != nil {
		l.Head.Next = item
	}
	l.Head = item
	l.Length++
	if l.Tail == nil {
		l.Tail = l.Head
	}
	return l.Head
}

func (l *list) PushFront(val interface{}) *ListItem {
	item := &ListItem{Prev: nil, Value: val, Next: l.Tail}
	if l.Tail != nil {
		l.Tail.Prev = item
	}
	l.Tail = item
	l.Length++
	if l.Head == nil {
		l.Head = l.Tail
	}
	return l.Tail
}

func (l *list) Remove(item *ListItem) {
	if item.Prev != nil {
		item.Prev.Next = item.Next
	} else {
		l.Tail = item.Next
	}

	if item.Next != nil {
		item.Next.Prev = item.Prev
	} else {
		l.Head = item.Prev
	}
	l.Length--
}

func (l *list) MoveToFront(item *ListItem) {
	l.Remove(item)
	l.PushFront(item.Value)
}

func NewList() List {
	return new(list)
}
