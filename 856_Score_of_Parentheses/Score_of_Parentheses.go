package scoreofparentheses

import "fmt"

//(()()())
func ScoreOfParentheses(s string) int {
	stack := []rune{}
	score := 0
	inner := 0
	for _, val := range s {
		if val == 0x28 {
			stack = append(stack, val)
			inner = 0
		} else {
			stack = stack[1:]
			if len(stack) == 0 {
				if inner != 0 {
					score = score + inner*2
				} else {
					score++
				}

			} else {
				inner++
			}
		}
		fmt.Println(stack)
	}
	return score
}
