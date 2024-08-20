package kidswiththegreatestnumberofcandies

func KidsWithCandies(candies []int, extraCandies int) []bool {
	result := make([]bool, len(candies))
	max := 0
	for _, val := range candies {
		if val > max {
			max = val
		}
	}

	for i, val := range candies {
		if val+extraCandies < max {
			result[i] = false
		} else {
			result[i] = true
		}
	}

	return result
}
