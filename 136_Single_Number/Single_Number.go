package singlenumber

func SingleNumber(nums []int) int {
	//第一版
	// count := 0
	// for len(nums) != 1 {
	// 	tg := nums[count]
	// 	for i := count + 1; i < len(nums); i++ {

	// 		if tg == nums[i] {
	// 			nums = append(nums[:i], nums[i+1:]...)
	// 			nums = append(nums[:count], nums[count+1:]...)
	// 			count = 0
	// 			break
	// 		} else {
	// 			if i == len(nums)-1 {
	// 				count++
	// 			}
	// 		}
	// 	}

	// }
	// return nums[0]

	//第二版
	result := 0
	for _, data := range nums {
		result ^= data
	}
	return result
}
