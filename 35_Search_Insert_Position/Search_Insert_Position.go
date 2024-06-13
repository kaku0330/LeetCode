package searchinsertposition

func SearchInsert(nums []int, target int) int {
	index := len(nums) / 2
	for index >= 0 && index < len(nums) {
		if nums[index] < target {
			index++
			if index < len(nums) && nums[index] > target {
				return index
			}
		} else if nums[index] == target {
			return index
		} else {
			index--
			if index > 0 && nums[index] < target {
				return index + 1
			}
		}
	}
	if index < 0 {
		return index + 1
	} else {
		return index
	}
}
