package maximumnumberofwordsfoundinsentences

import "strings"

func MostWordsFound(sentences []string) int {
	var check int
	for i := 0; i < len(sentences); i++ {
		if len((strings.Split(sentences[i], " "))) > check {
			check = len((strings.Split(sentences[i], " ")))
		}
	}
	return check
}
