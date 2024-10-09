package concatenationofarray

func GetConcatenation(nums []int) []int {
	for index := 0; index < len(nums)-index; index++ {
		nums = append(nums, nums[index])
	}

	return nums
}
