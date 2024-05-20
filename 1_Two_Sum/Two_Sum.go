package twosum

func TwoSum(nums []int, target int) []int {
	//第一版
	// type hashval struct {
	// 	value int
	// 	count int
	// }

	// numsMap := map[hashval]int{}

	// for i, num := range nums {
	// 	_, exist := numsMap[hashval{num, 0}]
	// 	if exist {
	// 		numsMap[hashval{num, 1}] = i
	// 	} else {
	// 		numsMap[hashval{num, 0}] = i
	// 	}
	// }

	// for i := 0; i < len(nums); i++ {
	// 	_, exist := numsMap[hashval{target - nums[i], 0}]
	// 	if exist {

	// 		if numsMap[hashval{nums[i], 0}] != numsMap[hashval{target - nums[i], 0}] {
	// 			return []int{numsMap[hashval{nums[i], 0}], numsMap[hashval{target - nums[i], 0}]}
	// 		} else {
	// 			_, doublecheck := numsMap[hashval{target - nums[i], 1}]
	// 			if doublecheck {
	// 				return []int{numsMap[hashval{nums[i], 0}], numsMap[hashval{target - nums[i], 1}]}
	// 			}
	// 		}
	// 	}
	// }

	// return []int{}

	//思路 用一個for迴圈跑如果map裡面沒有(targer-index的值)的話就加進map裡跑到有值才結束
	//第二版
	indexMap := make(map[int]int)
	for index, value := range nums {
		if val, exist := indexMap[target-value]; exist {
			return []int{val, index}
		}
		indexMap[value] = index
	}
	return []int{}
}
