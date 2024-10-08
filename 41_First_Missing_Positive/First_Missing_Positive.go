package firstmissingpositive

func FirstMissingPositive(nums []int) int {
	maps := make(map[int]bool, len(nums))

	for i := range nums {
		maps[nums[i]] = true
	}

	for i := 1; i < 2147483648; i++ {
		if !maps[i] {
			return i
		}
	}

	return 0
}
