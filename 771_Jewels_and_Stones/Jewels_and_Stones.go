package jewelsandstones

func NumJewelsInStones(jewels string, stones string) int {
	jewelmap := make(map[rune]byte, len(jewels))
	count := 0
	for _, val := range jewels {
		jewelmap[val] = 1
	}

	for _, val := range stones {
		if _, ok := jewelmap[val]; ok {
			count++
		}
	}
	return count
}
