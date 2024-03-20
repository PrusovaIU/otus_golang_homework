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
	List
	front *ListItem
	back  *ListItem
	len   int
}

func (l list) Len() int {
	return l.len
}

func (l list) Front() *ListItem {
	return l.front
}

func (l list) Back() *ListItem {
	return l.back
}

func (l *list) PushFront(v interface{}) *ListItem {
	newItem := ListItem{Value: v, Next: l.front, Prev: nil}
	if l.front != nil {
		l.front.Prev = &newItem
	}
	l.front = &newItem
	if l.back == nil {
		l.back = l.front
	}
	l.len++
	return l.front
}

func (l *list) PushBack(v interface{}) *ListItem {
	newItem := ListItem{Value: v, Next: nil, Prev: l.back}
	if l.back != nil {
		l.back.Next = &newItem
	}
	l.back = &newItem
	if l.front == nil {
		l.front = l.back
	}
	l.len++
	return l.back
}

func (l *list) sieze(i *ListItem) {
	if i.Prev != nil {
		i.Prev.Next = i.Next
	}
	if i.Next != nil {
		i.Next.Prev = i.Prev
	}

}

func (l *list) Remove(i *ListItem) {
	if l.front == i {
		l.front = i.Next
	} else if l.back == i {
		l.back = i.Prev
	}
	l.sieze(i)
	l.len--
}

func (l *list) MoveToFront(i *ListItem) {
	if l.front != i {
		l.sieze(i)
		i.Next = l.front
		i.Prev = nil
		if l.front != nil {
			l.front.Prev = i
		}
		l.front = i
	}
}

func NewList() List {
	return new(list)
}
