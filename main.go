package main

// import (
// 	twoSum "LeetCode/1_Two_Sum"
// 	"fmt"
// )
// import (
// 	validParentheses "LeetCode/20_Valid_Parentheses"
// 	"fmt"
// )
// import (
// removeElement "LeetCode/27_Remove_Element"
// 	"fmt"
// )
// import (
// 	besttimetobuyandsellstock "LeetCode/121_Best_Time_to_Buy_and_Sell_Stock"
// 	"fmt"
// )
// import (
// 	singlenumber "LeetCode/136_Single_Number"
// 	"fmt"
// )
// import (
// 	containsduplicate "LeetCode/217_Contains_Duplicate"
// 	"fmt"
// )
// import (
// palindromeLinkedList "LeetCode/234_Palindrome_Linked_List"
// 	"fmt"
// )
import (
	validanagram "LeetCode/242_Valid_Anagram"
	"fmt"
)

// import (
// 	finalvalueofvariableafterperformingoperations "LeetCode/2011_Final_Value_of_Variable_After_Performing_Operations"
// 	"fmt"
// )
// import (
// 	maximumnumberofwordsfoundinsentences "LeetCode/2114_Maximum_Number_of_Words_Found_in_Sentences"
// 	"fmt"
// )

func main() {
	//1題
	// nums := []int{3, 3}
	// fmt.Println(twoSum.TwoSum(nums, 6))

	//20題
	// s := "))"
	// fmt.Println(validParentheses.IsValid(s))

	//27題  題目怪怪的先不寫
	// num := []int{1, 2, 3, 4, 4}
	// fmt.Println(removeElement.RemoveElement(num, 1))

	//121題
	// prices := []int{7, 1, 5, 3, 6, 4}
	// fmt.Println(besttimetobuyandsellstock.MaxProfit(prices))

	//136題
	// num := []int{1, 2, 4, 1, 2}
	// fmt.Println(singlenumber.SingleNumber(num))

	//217題
	// nums := []int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2}
	// fmt.Println(containsduplicate.ContainsDuplicate(nums))
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

	//242題
	s, t := "anagram", "nagaram"
	fmt.Println(validanagram.IsAnagram(s, t))

	//2011題
	// operations := []string{"X++", "++X", "--X", "X--"}
	// fmt.Println(finalvalueofvariableafterperformingoperations.FinalValueAfterOperations(operations))

	//2114題
	// sentences := []string{"alice and bob love leetcode", "i think so too", "this is great thanks very much"}
	// fmt.Println(maximumnumberofwordsfoundinsentences.MostWordsFound(sentences))
}
