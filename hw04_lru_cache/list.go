package hw04lrucache

import "fmt"

// List - интерфейс.
type List interface {
	Len() int
	Front() *ListItem
	Back() *ListItem
	PushFront(v interface{}) *ListItem
	PushBack(v interface{}) *ListItem
	Remove(i *ListItem)
	MoveToFront(i *ListItem)
}

// NewList - конструктор для создания объекта интерфейса.
func NewList() List {
	return newList()
}

// ListItem - структура ноды.
type ListItem struct {
	Value interface{}
	Next  *ListItem
	Prev  *ListItem
}

// NewListItem - Конструктор структуры ListItem.
func NewListItem(value interface{}) *ListItem {
	return &ListItem{
		Value: value,
		Next:  nil,
		Prev:  nil,
	}
}

// list - структура двусвязного списка.
type list struct {
	Size int
	Head *ListItem
	Tail *ListItem
}

// newList - Конструктор для list.
func newList() *list {
	return &list{
		Size: 0,
		Head: nil,
		Tail: nil,
	}
}

// Len - возвращение длины списка.
func (l *list) Len() int {
	return l.Size
}

// Front - получение первой ноды списка.
func (l *list) Front() *ListItem {
	// не понимаю как протестировать nil в ответе, по заданию сигнатура дана такая
	if l.Head == nil {
		fmt.Println("Список пуст")
		return nil
	}
	return l.Head
}

// Back - получение последней ноды списка.
func (l *list) Back() *ListItem {
	// не понимаю как протестировать nil в ответе, по заданию сигнатура дана такая
	if l.Tail == nil {
		fmt.Println("Список пуст")
		return nil
	}
	return l.Tail
}

// PushFront - добавление новой ноды в начало списка.
func (l *list) PushFront(v interface{}) *ListItem {
	newNode := NewListItem(v)
	if l.Head == nil {
		l.Head = newNode
		l.Tail = newNode
	} else {
		newNode.Next = l.Head
		l.Head.Prev = newNode
		l.Head = newNode
	}
	l.Size++

	return newNode
}

// PushBack - добавление новой ноды в конец списка.
func (l *list) PushBack(v interface{}) *ListItem {
	newNode := NewListItem(v)
	if l.Head == nil {
		l.Head = newNode
		l.Tail = newNode
	} else {
		currentNode := l.Head
		for currentNode.Next != nil {
			currentNode = currentNode.Next
		}
		newNode.Prev = currentNode
		currentNode.Next = newNode
		l.Tail = newNode
	}
	l.Size++
	return newNode
}

// Remove - удаление ноды.
func (l *list) Remove(i *ListItem) {
	// метод не удаляет от Prev в начало
	// не могу понять, как сделать
	node := l.Head
	if node.Value == i.Value {
		node = node.Next
		node.Prev = node.Prev.Prev
		l.Head = node
	} else {
		for node != nil {
			if node.Next.Value == i.Value {
				node.Next = node.Next.Next
				if node.Next != nil {
					node.Prev = node.Next.Prev
				}
				break
			} else {
				node = node.Next
			}
		}
	}
	l.Size--
}

// MoveToFront - перемещение ноды в начало списка.
func (l *list) MoveToFront(i *ListItem) {
	// сделал через создание новой ноды и перекладки в нее значения
	// если передавать ноду и ее использовать, то закольцовывается l.Next
	// как победить не понял
	newNode := NewListItem(i.Value)
	newNode.Next = l.Head
	l.Head.Prev = newNode
	l.Head = newNode
	node := l.Head
	for node != nil {
		if node.Next.Value == i.Value {
			node.Next = node.Next.Next
			if node.Next != nil {
				node.Prev = node.Next.Prev
			}
			break
		} else {
			node = node.Next
		}
	}
}

// Print - печать (не треуется по заданию)
func (l *list) Print() {
	if l.Head == nil {
		fmt.Println("Список пуст")
	}
	node := l.Head
	for node != nil {
		fmt.Println(node.Value)
		node = node.Next
	}
}
