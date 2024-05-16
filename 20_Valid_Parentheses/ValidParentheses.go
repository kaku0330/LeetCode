package validParentheses

func IsValid(s string) bool {
	//第一次
	// var first []rune
	// if len(s)%2 != 0 {
	// 	return false
	// }
	// for _, element := range s {
	// 	if element == 0x28 || element == 0x5B || element == 0x7B {
	// 		first = append(first, element)
	// 	} else if element == 0x29 || element == 0x5D || element == 0x7D {
	// 		if len(first) != 0 {
	// 			if first[len(first)-1] == 0x28 {
	// 				if element == 0x29 {
	// 					first = first[0 : len(first)-1]
	// 				} else {
	// 					return false
	// 				}
	// 			} else if first[len(first)-1] == 0x5B {
	// 				if element == 0x5D {
	// 					first = first[0 : len(first)-1]
	// 				} else {
	// 					return false
	// 				}
	// 			} else if first[len(first)-1] == 0x7B {
	// 				if element == 0x7D {
	// 					first = first[0 : len(first)-1]
	// 				} else {
	// 					return false
	// 				}
	// 			}
	// 		} else {
	// 			return false
	// 		}
	// 	}
	// }
	// if len(first) == 0 {
	// 	return true
	// } else {
	// 	return false
	// }

	//第二次
	var first []rune
	check := map[rune]rune{
		0x29: 0x28,
		0x5D: 0x5B,
		0x7D: 0x7B,
	}

	if len(s)%2 != 0 {
		return false
	}

	for _, element := range s {
		switch element {
		case 0x28, 0x5B, 0x7B:
			first = append(first, element)
		case 0x29, 0x5D, 0x7D:
			if len(first) != 0 {
				if first[len(first)-1] != check[element] {
					return false
				} else {
					first = first[0 : len(first)-1]
				}
			} else {
				return false
			}
		}
	}
	return len(first) == 0
}
