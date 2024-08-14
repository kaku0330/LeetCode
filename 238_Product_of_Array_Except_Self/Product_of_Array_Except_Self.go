package productofarrayexceptself

import "fmt"

func ProductExceptSelf(nums []int) []int {
	sum := 0

	for _, val := range nums {
		if val == 0 {
			continue
		}
		if sum == 0 {
			sum = val
		} else {
			sum *= val
		}
	}
	fmt.Println(sum)
	return nums
}
