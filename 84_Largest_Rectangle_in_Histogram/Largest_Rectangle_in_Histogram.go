package largestrectangleinhistogram

func LargestRectangleArea(heights []int) int {
	highest := 0
	if len(heights) < 2 {
		return heights[0]
	}
	for i := 0; i < len(heights)-1; i++ {
		cache := 0
		if heights[i] < heights[i+1] {
			if heights[i] == 0 {
				cache = heights[i+1]
			} else {
				cache = heights[i] * 2
			}
		} else {
			if heights[i+1] == 0 {
				cache = heights[i]
			} else {
				cache = heights[i+1] * 2
			}
		}
		if cache > highest {
			highest = cache
		}
	}
	return highest
}
