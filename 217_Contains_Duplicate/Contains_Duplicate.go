package containsduplicate

func ContainsDuplicate(nums []int) bool {
	indexMap := make(map[int]int)
	for index, value := range nums {
		if _, exist := indexMap[value]; exist {
			return true
		}
		indexMap[value] = index
	}
	return false
}
