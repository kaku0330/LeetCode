package findgreatestcommondivisorofarray

func FindGCD(nums []int) int {
	max, min := nums[0], nums[0]
	for _, num := range nums {
		if num > max {
			max = num
		}
		if num < min {
			min = num
		}
	}

	for min != 0 {
		max, min = min, max%min
	}

	return max
}
