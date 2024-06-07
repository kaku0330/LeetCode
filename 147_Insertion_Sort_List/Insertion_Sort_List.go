package insertionsortlist

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

// 4 2 1 3  如果index1 > index2   index2指向1 1指向2指向的地方
func InsertionSortList(head *ListNode) *ListNode {

	// i := 0
	// for head.Next != nil {

	// 	if len(num) == 0 {
	// 		num = append(num, *head)
	// 	} else if len(num) == 1 {
	// 		if num[0].Val > *&head.Val {

	// 		}
	// 	}
	// 	head = head.Next
	// }
	// fmt.Println(num)
	return head
}
