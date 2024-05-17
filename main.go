package main

import (
	// validParentheses "LeetCode/20_Valid_Parentheses"
	palindromeLinkedList "LeetCode/234_Palindrome_Linked_List"
	"fmt"
)

func main() {
	// s := "))"
	// fmt.Println(validParentheses.IsValid(s))

	e := palindromeLinkedList.ListNode{
		Val:  1,
		Next: nil,
	}
	d := palindromeLinkedList.ListNode{
		Val:  1,
		Next: &e,
	}

	c := palindromeLinkedList.ListNode{
		Val:  4,
		Next: &d,
	}
	b := palindromeLinkedList.ListNode{
		Val:  1,
		Next: &c,
	}
	a := palindromeLinkedList.ListNode{
		Val:  1,
		Next: &b,
	}
	fmt.Println(palindromeLinkedList.IsPalindrome(&a))
}
