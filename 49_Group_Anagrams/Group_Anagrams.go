package groupanagrams

import "fmt"

//{"eat", "tea", "tan", "ate", "nat", "bat"}
func GroupAnagrams(strs []string) [][]string {
	count := 1
	indexArr := [][]string{}
	for len(strs) != 0 {
		if len(strs) >= count {
			fmt.Println(strs[:count])
			indexArr = append(indexArr, strs[:count])
			strs = strs[count:]
		} else {
			indexArr = append(indexArr, strs[:len(strs)-1])
			strs = strs[count:count]
		}
		count++
	}

	return indexArr
}
