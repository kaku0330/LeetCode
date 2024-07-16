package permutationdifferencebetweentwostrings

func FindPermutationDifference(s string, t string) int {
	stringmap := make(map[rune]int, len(s))
	count := 0
	for i, val := range s {
		stringmap[val] = i
	}

	for i, val := range t {
		num := stringmap[val] - i
		if num < 0 {
			num *= -1
		}
		count += num
	}

	return count
}
