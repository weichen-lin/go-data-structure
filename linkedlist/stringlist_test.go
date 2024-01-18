package linkedList

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_StringToLinkedList(t *testing.T) {
	var str string = "abcdefghijklmnopqrstuvwxyz"
	l1 := StringToLinkedList(str)
	l2 := StringToLinkedListV2(str)

	require.Equal(t, len(str), l1.Length)
	require.Equal(t, len(str), l2.Length)

	currentNode1 := l1.Head
	currentNode2 := l2.Head

	for i := 0; i < len(str); i++ {
		require.Equal(t, currentNode1.Value, currentNode2.Value)
		currentNode1 = currentNode1.Next
		currentNode2 = currentNode2.Next
	}
}

func Test_IsPalindrome(t *testing.T) {
	str1 := "abcba"
	str2 := "abccba"
	str3 := "abccdba"
	str4 := ""

	require.Equal(t, true, IsPalindrome(str1))
	require.Equal(t, true, IsPalindrome(str2))
	require.Equal(t, false, IsPalindrome(str3))
	require.Equal(t, true, IsPalindrome(str4))
}
