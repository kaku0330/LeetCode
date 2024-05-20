package twosum

func TwoSum(nums []int, target int) []int {

	type hashval struct {
		value int
		count int
	}

	numsMap := map[hashval]int{}

	for i, num := range nums {
		_, exist := numsMap[hashval{num, 0}]
		if exist {
			numsMap[hashval{num, 1}] = i
		} else {
			numsMap[hashval{num, 0}] = i
		}

	}
	for i := 0; i < len(nums); i++ {
		_, exist := numsMap[hashval{target - nums[i], 0}]
		if exist {

			if numsMap[hashval{nums[i], 0}] != numsMap[hashval{target - nums[i], 0}] {
				return []int{numsMap[hashval{nums[i], 0}], numsMap[hashval{target - nums[i], 0}]}
			} else {
				_, doublecheck := numsMap[hashval{target - nums[i], 1}]
				if doublecheck {
					return []int{numsMap[hashval{nums[i], 0}], numsMap[hashval{target - nums[i], 1}]}
				}
			}
		}
	}

	return []int{}
}
