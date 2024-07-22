package reversewordsinastring

import (
	"strings"
)

func ReverseWords(s string) string {
	//版本1  5ms Beats 23.60%  3.52MB Beats 57.14%
	// splitS := strings.Split(s, " ")

	// for i := 0; i < len(splitS); {
	// 	if splitS[i] == "" {
	// 		splitS = append(splitS[:i], splitS[i+1:]...)
	// 	} else {
	// 		i++
	// 	}
	// }

	// for i := 0; i < len(splitS)/2; i++ {
	// 	splitS[i], splitS[len(splitS)-1-i] = splitS[len(splitS)-1-i], splitS[i]
	// }

	// return strings.Join(splitS, " ")

	//版本2 0ms Beats 100.00%  3.55MB Beats 57.14%
	splitS := strings.Split(s, " ")
	for i := 0; i < len(splitS)/2; i++ {
		if splitS[i] == "" {
			splitS = append(splitS[:i], splitS[i+1:]...)
			i--
		} else if splitS[len(splitS)-1-i] == "" {
			splitS = append(splitS[:len(splitS)-1-i], splitS[len(splitS)-1-i+1:]...)
			i--
		} else {
			splitS[i], splitS[len(splitS)-1-i] = splitS[len(splitS)-1-i], splitS[i]
		}
	}

	if len(splitS)%2 == 1 && splitS[len(splitS)/2] == "" {
		splitS = append(splitS[:len(splitS)/2], splitS[len(splitS)/2+1:]...)
	}
	return strings.Join(splitS, " ")
}
