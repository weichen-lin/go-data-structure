package linkedList

type StringNode struct {
	Value rune
	Next  *StringNode
}

type StringLinkedList struct {
	Head   *StringNode
	Length int
}

func StringToLinkedList(str string) *StringLinkedList {
	var l StringLinkedList
	for _, c := range str {
		if( l.Head == nil ) {
			l.Head = &StringNode{Value: c}
		} else {
			currentNode := l.Head
			for currentNode.Next != nil {
				currentNode = currentNode.Next
			}
			currentNode.Next = &StringNode{Value: c}
		}
	}
	return &l
}

func isPalindrome(l *StringLinkedList) bool {
	if l.Head == nil  {
		return true
	}

	if l.Head.Next == nil {
		return true
	}

	slow, fast := l.Head, l.Head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	
	var prev *StringNode	
	current := slow
	for current != nil {
		nextTemp := current.Next
		current.Next = prev
		prev = current
		current = nextTemp
	}

	// 比较前半部分和反转后的后半部分是否相等
	left, right := head, prev
	for right != nil {
		if left.Val != right.Val {
			return false
		}
		left = left.Next
		right = right.Next
	}

	return true
}
