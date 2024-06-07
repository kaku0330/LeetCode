package main

// import (
// 	twoSum "LeetCode/1_Two_Sum"
// 	"fmt"
// )
import (
	romantointeger "LeetCode/13_Roman_To_Integer"
	"fmt"
)

// import (
// 	validParentheses "LeetCode/20_Valid_Parentheses"
// 	"fmt"
// )
// import (
// removeElement "LeetCode/27_Remove_Element"
// 	"fmt"
// )
// import (
// 	groupanagrams "LeetCode/49_Group_Anagrams"
// 	"fmt"
// )
// import (
// 	lengthOfLastWord "LeetCode/58_Length_of_Last_Word"
// 	"fmt"
// )
// import (
// 	simplifypath "LeetCode/71_Simplify_Path"
// 	"fmt"
// )
// import (
// 	largestrectangleinhistogram "LeetCode/84_Largest_Rectangle_in_Histogram"
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
// 	insertionSortList "LeetCode/147_Insertion_Sort_List"
// 	"fmt"
// )

// import (
// 	minstack "LeetCode/155_Min_Stack"
// 	"fmt"
// )
// import (
// 	containsduplicate "LeetCode/217_Contains_Duplicate"
// 	"fmt"
// )
// import (
// 	implementstackusingqueues "LeetCode/225_Implement_Stack_using_Queues"
// 	"fmt"
// )
// import (
// 	implementqueueusingstacks "LeetCode/232_Implement_Queue_using_Stacks"
// 	"fmt"
// )
// import (
// palindromeLinkedList "LeetCode/234_Palindrome_Linked_List"
// 	"fmt"
// )
// import (
// 	validanagram "LeetCode/242_Valid_Anagram"
// 	"fmt"
// )
// import (
// 	decodestring "LeetCode/394_Decode_String"
// 	"fmt"
// )
// import (
// 	baseballgame "LeetCode/682_Baseball_Game"
// 	"fmt"
// )
// import (
// 	scoreOfParentheses "LeetCode/856_Score_of_Parentheses"
// 	"fmt"
// )

// import (
// 	numberofstudentsunabletoeatlunch "LeetCode/1700_Number_of_Students_Unable_to_Eat_Lunch"
// 	"fmt"
// )
// import (
// 	finalvalueofvariableafterperformingoperations "LeetCode/2011_Final_Value_of_Variable_After_Performing_Operations"
// 	"fmt"
// )
// import (
// 	maximumnumberofwordsfoundinsentences "LeetCode/2114_Maximum_Number_of_Words_Found_in_Sentences"
// 	"fmt"
// )
// import (
// 	sortvowelsinastring "LeetCode/2785_Sort_Vowels_in_a_String"
// 	"fmt"
// )
// import (
// 	scoreofastring "LeetCode/3110_Score_of_a_String"
// 	"fmt"
// )

func main() {
	//1題
	// nums := []int{3, 3}
	// fmt.Println(twoSum.TwoSum(nums, 6))

	//13題
	s := "MDCCCLXXXIV"
	fmt.Println(romantointeger.RomanToInt(s))

	//20題
	// s := "))"
	// fmt.Println(validParentheses.IsValid(s))

	//27題  題目怪怪的先不寫
	// num := []int{1, 2, 3, 4, 4}
	// fmt.Println(removeElement.RemoveElement(num, 1))

	//49題
	// strs := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	// fmt.Println(groupanagrams.GroupAnagrams(strs))

	//58題
	// s := "   fly me   to   the moon  "
	// fmt.Println(lengthOfLastWord.LengthOfLastWord(s))

	//71題
	// path := "/../"
	// fmt.Println(simplifypath.SimplifyPath(path))

	//84題
	// heights := []int{2, 1, 5, 6, 2, 3}
	// fmt.Println(largestrectangleinhistogram.LargestRectangleArea(heights))

	//121題
	// prices := []int{7, 1, 5, 3, 6, 4}
	// fmt.Println(besttimetobuyandsellstock.MaxProfit(prices))

	//136題
	// num := []int{1, 2, 4, 1, 2}
	// fmt.Println(singlenumber.SingleNumber(num))

	//147題
	// e := insertionSortList.ListNode{
	// Val:  1,
	// Next: nil,
	// }
	// d := insertionSortList.ListNode{
	// 	Val:  3,
	// 	Next: nil,
	// }

	// c := insertionSortList.ListNode{
	// 	Val:  1,
	// 	Next: &d,
	// }
	// b := insertionSortList.ListNode{
	// 	Val:  2,
	// 	Next: &c,
	// }
	// a := insertionSortList.ListNode{
	// 	Val:  4,
	// 	Next: &b,
	// }
	// fmt.Println(insertionSortList.InsertionSortList(&a))
	//155題
	// minStack := minstack.Constructor()
	// minStack.Push(2147483646)
	// minStack.Push(2147483646)
	// minStack.Push(2147483647)
	// fmt.Println(minStack.Top())
	// minStack.Pop()
	// fmt.Println(minStack.GetMin())
	// minStack.Pop()
	// fmt.Println(minStack.GetMin())
	// minStack.Pop()
	// minStack.Push(2147483647)
	// fmt.Println(minStack.Top())
	// fmt.Println(minStack.GetMin())
	// minStack.Push(-2147483647)
	// fmt.Println(minStack.Top())
	// fmt.Println(minStack.GetMin())
	// minStack.Pop()
	// fmt.Println(minStack.GetMin())

	//217題
	// nums := []int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2}
	// fmt.Println(containsduplicate.ContainsDuplicate(nums))

	//225題
	// minStack := implementstackusingqueues.Constructor()
	// minStack.Push(1)
	// minStack.Push(2)
	// fmt.Println(minStack.Top())
	// fmt.Println(minStack.Pop())
	// fmt.Println(minStack.Empty())

	//232題
	// minStack := implementqueueusingstacks.Constructor()
	// minStack.Push(1)
	// minStack.Push(2)
	// fmt.Println(minStack.Peek())
	// fmt.Println(minStack.Pop())
	// fmt.Println(minStack.Empty())

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
	// s, t := "anagram", "nagaram"
	// fmt.Println(validanagram.IsAnagram(s, t))

	//394題
	// s := "3[a]2[bc]"
	// fmt.Println(decodestring.DecodeString(s))

	//682題
	// operations := []string{"5", "2", "C", "D", "+"}
	// fmt.Println(baseballgame.CalPoints(operations))

	//859題
	// s := "(()(()))"
	// fmt.Println(scoreOfParentheses.ScoreOfParentheses(s))

	//1700題
	// students, sandwiches := []int{1, 1, 0, 0}, []int{0, 1, 0, 1}
	// students, sandwiches := []int{1, 1, 1, 0, 0, 1}, []int{1, 0, 0, 0, 1, 1}
	// fmt.Println(numberofstudentsunabletoeatlunch.CountStudents(students, sandwiches))

	//2011題
	// operations := []string{"X++", "++X", "--X", "X--"}
	// fmt.Println(finalvalueofvariableafterperformingoperations.FinalValueAfterOperations(operations))

	//2114題
	// sentences := []string{"alice and bob love leetcode", "i think so too", "this is great thanks very much"}
	// fmt.Println(maximumnumberofwordsfoundinsentences.MostWordsFound(sentences))

	//2785題
	// s := "lEetcOde"
	// fmt.Println(sortvowelsinastring.SortVowels(s))

	//3110題
	// s := "hello"
	// fmt.Println(scoreofastring.ScoreOfString(s))
}
