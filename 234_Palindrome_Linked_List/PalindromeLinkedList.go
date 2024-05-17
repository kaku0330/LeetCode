package palindromelinkedlist

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func IsPalindrome(head *ListNode) bool {
	length := 1

	var value []int
	for head.Next != nil {
		length++
		value = append(value, head.Val)
		head = head.Next
	}
	value = append(value, head.Val)
	if length == 1 {
		return true
	}
	left, right := 0, length-1
	if length%2 == 0 {
		for right-left != -1 {
			if value[left] != value[right] {
				return false
			}
			right--
			left++
		}
	} else {
		for right != left {
			if value[left] != value[right] {
				return false
			}
			right--
			left++
		}
	}
	return true
}
