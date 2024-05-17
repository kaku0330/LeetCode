package main

// validParentheses "LeetCode/20_Valid_Parentheses"
// removeElement "LeetCode/27_Remove_Element"
import (
	singlenumber "LeetCode/136_Single_Number"
	"fmt"
)

// palindromeLinkedList "LeetCode/234_Palindrome_Linked_List"

func main() {
	//20題
	// s := "))"
	// fmt.Println(validParentheses.IsValid(s))

	//27題  題目怪怪的先不寫
	// num := []int{1, 2, 3, 4, 4}
	// fmt.Println(removeElement.RemoveElement(num, 1))

	//136題
	num := []int{1, 2, 4, 1, 2}
	fmt.Println(singlenumber.SingleNumber(num))

	//234題
	// e := palindromeLinkedList.ListNode{
	// 	Val:  1,
	// 	Next: nil,
	// }
	// d := palindromeLinkedList.ListNode{
	// 	Val:  1,
	// 	Next: &e,
	// }

	// c := palindromeLinkedList.ListNode{
	// 	Val:  4,
	// 	Next: &d,
	// }
	// b := palindromeLinkedList.ListNode{
	// 	Val:  1,
	// 	Next: &c,
	// }
	// a := palindromeLinkedList.ListNode{
	// 	Val:  1,
	// 	Next: &b,
	// }
	// fmt.Println(palindromeLinkedList.IsPalindrome(&a))
}
