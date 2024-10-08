package movezeroes

func MoveZeroes(nums []int) []int {
	// pointer1 := 0
	// pointer2 := len(nums) - 1
	// fmt.Println(pointer1, pointer2)
	// for pointer1 <= pointer2 {
	// 	fmt.Println(pointer1, pointer2)
	// 	if nums[pointer2] == 0 && nums[pointer2] != len(nums)-1 {
	// 		nums = append(nums[:pointer2], nums[pointer2+1:]...)
	// 		nums = append(nums, 0)
	// 		pointer2--
	// 	} else {
	// 		pointer2--
	// 	}
	// 	if nums[pointer1] == 0 {
	// 		nums = append(nums[:pointer1], nums[pointer1+1:]...)
	// 		nums = append(nums, 0)
	// 		pointer2--
	// 	} else {
	// 		pointer1++
	// 	}
	// }

	// p1, p2 := 0, 1
	// for p2 < len(nums) {
	// 	if nums[p1] == 0 && nums[p2] != 0 {
	// 		nums[p1], nums[p2] = nums[p2], nums[p1]
	// 	}

	// 	if nums[p1] != 0 {
	// 		p1++
	// 	}

	// 	if nums[p2] == 0 || p2 <= p1 {
	// 		p2++
	// 	}
	// }
	// last := 0
	// for _, val := range nums {
	// 	if val != 0 {
	// 		nums[last] = val
	// 		last++
	// 	}
	// }

	// for ; last < len(nums); last++ {
	// 	nums[last] = 0
	// }

	last := 0
	for i := range nums {
		if nums[i] != 0 {
			nums[last], nums[i] = nums[i], nums[last]
			last++
		}
	}

	return nums
}
