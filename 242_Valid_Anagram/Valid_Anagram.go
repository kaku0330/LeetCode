package validanagram

func IsAnagram(s string, t string) bool {
	//第一版
	// sMap := make(map[string]int, len(s))
	// tMap := make(map[string]int, len(t))
	// if len(s) != len(t) {
	// 	return false
	// }
	// for i := 0; i < len(s); i++ {
	// 	_, exist := sMap[string(s[i])]
	// 	if exist {
	// 		sMap[string(s[i])] = sMap[string(s[i])] + 1
	// 	} else {
	// 		sMap[string(s[i])] = 1
	// 	}
	// 	_, exist = tMap[string(t[i])]
	// 	if exist {
	// 		tMap[string(t[i])] = tMap[string(t[i])] + 1
	// 	} else {
	// 		tMap[string(t[i])] = 1
	// 	}

	// }
	// for index, value := range sMap {
	// 	if tMap[index] != value {
	// 		return false
	// 	}
	// }
	// return true

	//第二版
	if len(s) != len(t) {
		return false
	}
	sMap := make(map[rune]int, len(s))
	for _, val := range s {
		sMap[val] = sMap[val] + 1
	}
	for _, val := range t {
		sMap[val] = sMap[val] - 1
		if sMap[val] < 0 {
			return false
		}
	}
	return true

}
