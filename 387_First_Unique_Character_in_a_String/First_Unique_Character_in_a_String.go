package firstuniquecharacterinastring

func FirstUniqChar(s string) int {
	hash := [26]int{}

	for _, val := range s {
		hash[val-'a'] = hash[val-'a'] + 1
	}

	for i, val := range s {
		if hash[val-'a'] == 1 {
			return i
		}
	}

	return -1
}
