package romantointeger

func RomanToInt(s string) int {
	romanNum := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}
	total := 0
	for i := 0; i < len(s); i++ {
		if i == len(s)-1 {
			total = total + romanNum[s[i]]
			break
		}
		if s[i] == 73 && (s[i+1] == 88 || s[i+1] == 86) {
			total = total + romanNum[s[i+1]] - romanNum[s[i]]
			i++
		} else if s[i] == 88 && (s[i+1] == 76 || s[i+1] == 67) {
			total = total + romanNum[s[i+1]] - romanNum[s[i]]
			i++
		} else if s[i] == 67 && (s[i+1] == 68 || s[i+1] == 77) {
			total = total + romanNum[s[i+1]] - romanNum[s[i]]
			i++
		} else {
			total = total + romanNum[s[i]]
		}
	}

	return total
}
