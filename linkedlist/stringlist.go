package linkedList

type StringNode struct {
	Value rune
	Next  *StringNode
}

type StringLinkedList struct {
	Head   *StringNode
	Length int
}

// 將字串轉換成鏈表，時間複雜度為 O(n2)
func StringToLinkedList(str string) *StringLinkedList {
	var l StringLinkedList
	for _, c := range str {
		if l.Head == nil {
			l.Head = &StringNode{Value: c}
		} else {
			currentNode := l.Head
			for currentNode.Next != nil {
				currentNode = currentNode.Next
			}
			currentNode.Next = &StringNode{Value: c}
		}
		l.Length++
	}
	return &l
}

// 將字串轉換成鏈表，時間複雜度為 O(n)
func StringToLinkedListV2(str string) *StringLinkedList {
	var l StringLinkedList

	var temp *StringNode

	for _, c := range str {
		if l.Head == nil {
			l.Head = &StringNode{Value: c}
			temp = l.Head
		} else {
			temp.Next = &StringNode{Value: c}

			// 臨時 Node 往後移一個
			temp = temp.Next
		}
		l.Length++
	}
	return &l
}

// 測試一個長度為 223 的字串轉換成鏈表
// V1 & V2 的時間複雜度差異大約為 5 倍

/*
核心思想：快慢指针，
- 快指针每次走两步，慢指针每次走一步。
 1. 奇數情況 : 當快指針走到鏈表尾部時，慢指針剛好走到鏈表中間。
 2. 偶數情況 : 當快指針走到鏈表尾部時，慢指針剛好走到鏈表中間的前一個節點。

- 找到練表中間的節點後，過程中慢指針所在節點之前的鏈表會被反轉。
- 比較慢指針前面被反轉的鏈表與慢指針開始的鏈表是否相同
*/
func IsPalindrome(s string) bool {
	l := StringToLinkedList(s)

	if l == nil || l.Head == nil {
		return true
	}

	slow, fast := l.Head, l.Head

	// 建立一個空 Node，用來存放慢指針所在節點之前的鏈表的比較起點 Node
	var nodeStartReverse *StringNode

	for fast != nil && fast.Next != nil {
		// 快指針每次走兩步
		fast = fast.Next.Next

		// 先拿到慢指針下一個的 Node
		nextNode := slow.Next

		// 將慢指針所在節點之前的鏈表反轉
		slow.Next = nodeStartReverse
		nodeStartReverse = slow
		slow = nextNode
	}

	// 偶數情況的話，快指針會走到鏈表外部，慢指針剛好走到鏈表中間的後一個節點
	//   O <- O <- O -> O -> O -> O  -> null
	//	          起點  慢                快
	// 慢指針不需要移動

	// 奇數情況，快指針會走到鏈表最後一位，慢指針剛好走到鏈表中間
	//   O <- O <- O -> O -> O
	//      起點    慢        快
	// 慢需要往後移一個，且慢指針所在節點之前的鏈表已經反轉過了
	if fast != nil {
		slow = slow.Next
	}

	// 比較反轉後的前半部分與後半部分
	for nodeStartReverse != nil && slow != nil {
		if nodeStartReverse.Value != slow.Value {
			return false
		}
		nodeStartReverse = nodeStartReverse.Next
		slow = slow.Next
	}

	return true
}
