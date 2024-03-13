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
	size      int
	firstItem *ListItem
	lastItem  *ListItem
}

func (list *list) Len() int {
	return list.size
}

func (list *list) Front() *ListItem {
	return list.firstItem
}

func (list *list) Back() *ListItem {
	return list.lastItem
}

func (list *list) PushFront(v interface{}) *ListItem {
	newListItem := &ListItem{v, list.Front(), nil}

	if list.firstItem != nil {
		list.firstItem.Prev = newListItem
	} else {
		list.lastItem = newListItem
	}

	list.firstItem = newListItem
	list.size++

	return newListItem
}

func (list *list) PushBack(v interface{}) *ListItem {
	newListItem := &ListItem{v, nil, list.Back()}

	if list.lastItem != nil {
		list.lastItem.Next = newListItem
	} else {
		list.firstItem = newListItem
	}

	list.lastItem = newListItem
	list.size++

	return newListItem
}

func (list *list) Remove(i *ListItem) {
	if i == nil {
		return
	}

	prevItem := i.Prev
	nextItem := i.Next

	if prevItem != nil {
		prevItem.Next = nextItem
	} else {
		list.firstItem = nextItem
	}

	if nextItem != nil {
		nextItem.Prev = prevItem
	} else {
		list.lastItem = prevItem
	}

	list.size--
}

func (list *list) MoveToFront(i *ListItem) {
	if i == nil {
		return
	}

	list.Remove(i)
	list.PushFront(i.Value)
}

func NewList() List {
	return new(list)
}
