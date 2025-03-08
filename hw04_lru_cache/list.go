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
	head   *ListItem
	tail   *ListItem
	length int
}

func (l list) Len() int {
	return l.length
}

func (l list) Front() *ListItem {
	return l.head
}

func (l list) Back() *ListItem {
	return l.tail
}

func (l *list) PushFront(v interface{}) *ListItem {
	item := &ListItem{}
	item.Value = v

	if l.head == nil {
		l.head = item
		l.tail = item
		item.Prev = nil
		item.Next = nil
	} else {
		l.head.Prev = item
		item.Next = l.head
		l.head = item
	}
	l.length++
	return item
}

func (l *list) PushBack(v interface{}) *ListItem {
	item := &ListItem{}
	item.Value = v
	if l.head == nil {
		l.head = item
		l.tail = item
		item.Prev = nil
		item.Next = nil
	} else {
		item.Prev = l.tail
		item.Next = nil
		l.tail.Next = item
		l.tail = item
	}
	l.length++

	return item
}

func (l *list) Remove(i *ListItem) {
	if l.length == 1 {
		l.tail = nil
		l.head = nil
		l.length--
		return
	}

	if i.Next == nil {
		l.tail = l.tail.Prev
		l.tail.Next = nil
	}

	if i.Prev == nil {
		l.head = l.head.Next
		l.head.Prev = nil
	}

	if i.Next != nil && i.Prev != nil {
		i.Prev.Next = i.Next
		i.Next.Prev = i.Prev
	}
	l.length--
}

func (l *list) MoveToFront(i *ListItem) {
	if i.Prev == nil {
		return
	}

	if i.Next == nil {
		l.tail = l.tail.Prev
		l.tail.Next = nil
	}

	l.head.Prev = i
	i.Next = l.head
	i.Prev = nil
	l.head = i
}

func NewList() List {
	return new(list)
}
