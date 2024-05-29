package lengthoflastword

import "strings"

func LengthOfLastWord(s string) int {
	last := ""
	for _, val := range strings.Split(s, " ") {
		if val != "" {
			last = val
		}
	}
	return len(last)
}
