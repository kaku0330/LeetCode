package numberofgoodpairs

func NumIdenticalPairs(nums []int) int {
	cached := make(map[int]int)
	total := 0
	for _, num := range nums {
		if _, ok := cached[num]; ok {
			cached[num]++
			total += cached[num]
		} else {
			cached[num] = 0
		}
	}
	return total
}
