package reversevowelsofastring

import (
	"strings"
)

func ReverseVowels(s string) string {
	//第一版
	// mapS := []byte{}
	// result := ""
	// for i := 0; i < len(s); i++ {
	// 	if s[i] == 'a' || s[i] == 'e' || s[i] == 'i' || s[i] == 'o' || s[i] == 'u' || s[i] == 'A' || s[i] == 'E' || s[i] == 'I' || s[i] == 'O' || s[i] == 'U' {
	// 		mapS = append(mapS, s[i])
	// 	}
	// }

	// for i, j := 0, len(mapS)-1; i < len(s); i++ {
	// 	if s[i] == 'a' || s[i] == 'e' || s[i] == 'i' || s[i] == 'o' || s[i] == 'u' || s[i] == 'A' || s[i] == 'E' || s[i] == 'I' || s[i] == 'O' || s[i] == 'U' {
	// 		result += string(mapS[j])
	// 		j--
	// 	} else {
	// 		result += string(s[i])
	// 	}
	// }

	// return result

	//第二版
	spliteS := strings.Split(s, "")
	for i, j := 0, len(s)-1; i < j; {
		if checkvowels(s[i]) && checkvowels(s[j]) {
			spliteS[i], spliteS[j] = spliteS[j], spliteS[i]
			i++
			j--
		}
		if !checkvowels(s[i]) {
			i++
		}
		if !checkvowels(s[j]) {
			j--
		}
	}

	return strings.Join(spliteS, "")
}

func checkvowels(s byte) bool {
	vowels := "aeiouAEIOU"
	for i := 0; i < len(vowels); i++ {
		if s == vowels[i] {
			return true
		}
	}
	return false
}
