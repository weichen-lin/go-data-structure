package linkedList

import (
	"errors"
	"sync"

	"github.com/google/uuid"
)

type Node struct {
	Value uuid.UUID
	Next  *Node
}

type SingleLinkedList struct {
	sync.Mutex
	Head   *Node
	Length int
}

func (l *SingleLinkedList) Append(value uuid.UUID) {
	l.Lock()
	defer l.Unlock()
	node := &Node{Value: value}

	if l.Head == nil {
		l.Head = node
	} else {
		currentNode := l.Head

		// 循環找到最後一個節點
		for currentNode.Next != nil {
			currentNode = currentNode.Next
		}

		currentNode.Next = node
	}

	l.Length++
}

func (l *SingleLinkedList) Delete(value uuid.UUID) error {
	l.Lock()
	defer l.Unlock()

	if l.Head == nil {
		return errors.New("this single linked list is empty")
	}

	if l.Head.Value == value {
		l.Head = l.Head.Next
		l.Length--
		return nil
	}

	currentNode := l.Head

	for currentNode.Next != nil {
		if currentNode.Next.Value == value {
			currentNode.Next = currentNode.Next.Next
			l.Length--
			return nil
		}

		currentNode = currentNode.Next
	}

	return errors.New("the value was not found in this single linked list")
}

func (l *SingleLinkedList) Prepend(value uuid.UUID) {
	l.Lock()
	defer l.Unlock()
	node := &Node{Value: value}

	if l.Head == nil {
		l.Head = node
	} else {
		node.Next = l.Head
		l.Head = node
	}

	l.Length++
}

func (l *SingleLinkedList) Search(value uuid.UUID) (int, error) {
	l.Lock()
	defer l.Unlock()

	if l.Head == nil {
		return 0, errors.New("this single linked list is empty")
	}

	currentNode := l.Head
	index := 0

	for currentNode != nil {
		if currentNode.Value == value {
			return index, nil
		}

		currentNode = currentNode.Next
		index++
	}

	return 0, errors.New("the value was not found in this single linked list")
}

func (l *SingleLinkedList) ValueOf(index int) (uuid.UUID, error) {
	l.Lock()
	defer l.Unlock()

	if index < 0 || index > l.Length {
		return uuid.Nil, errors.New("the index is out of range")
	}

	currentNode := l.Head
	currentIndex := 0

	for currentNode != nil {
		if currentIndex == index {
			return currentNode.Value, nil
		}

		currentNode = currentNode.Next
		currentIndex++
	}

	return uuid.Nil, errors.New("the index was not found in this single linked list")
}

func (l *SingleLinkedList) InsertBehindWithIndex(index int, value uuid.UUID) error {
	l.Lock()
	defer l.Unlock()

	if index < 0 {
		return errors.New("index can not be less than 0")
	}

	if l.Head == nil {
		return errors.New("this single linked list is empty")
	}

	if index == 0 {
		node := &Node{Value: value}
		node.Next = l.Head.Next
		l.Head.Next = node
		l.Length++
		return nil
	}

	node := &Node{Value: value}

	currentNode := l.Head
	currentIndex := 0

	for currentNode != nil {
		if currentIndex == index {
			node.Next = currentNode.Next
			currentNode.Next = node
			l.Length++
			return nil
		}

		currentNode = currentNode.Next
		currentIndex++
	}

	return errors.New("the index was not found in this single linked list, largest index must be equal to real position minus 1")
}

func (l *SingleLinkedList) Reverse() {
	l.Lock()
	defer l.Unlock()

	if l.Head == nil {
		return
	}

	if l.Head.Next == nil {
		return
	}

	var prev *Node
	current := l.Head
	
	// 思考: 一開始用 current.Next != nil 會有什麼問題?
	// Head -> node1 -> node2 -> nil (其實要反轉三次)
	// current.Next != nil 只會反轉兩次，實際上要反轉三次
	for current != nil {
		next := current.Next
		current.Next = prev
		prev = current
		current = next
	}
	l.Head = prev
}