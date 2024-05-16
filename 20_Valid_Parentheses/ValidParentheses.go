package validParentheses

func test(s string) bool {

	var first []rune
	if len(s)%2 != 0 {
		return false
	}
	for _, element := range s {
		if element == 0x28 || element == 0x5B || element == 0x7B {
			first = append(first, element)
		} else if element == 0x29 || element == 0x5D || element == 0x7D {
			if len(first) != 0 {
				if first[len(first)-1] == 0x28 {
					if element == 0x29 {
						first = first[0 : len(first)-1]
					} else {
						return false
					}
				} else if first[len(first)-1] == 0x5B {
					if element == 0x5D {
						first = first[0 : len(first)-1]
					} else {
						return false
					}
				} else if first[len(first)-1] == 0x7B {
					if element == 0x7D {
						first = first[0 : len(first)-1]
					} else {
						return false
					}
				}
			} else {
				return false
			}
		}
	}
	if len(first) == 0 {
		return true
	} else {
		return false
	}
}
