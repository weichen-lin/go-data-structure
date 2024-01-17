package main

import (
	"fmt"

	linkedList "github.com/weichen-lin/go-data/linkedlist"
)

func main() {
	a := "abcba"

	test := linkedList.IsPalindrome(a)
	fmt.Println(test)
}
