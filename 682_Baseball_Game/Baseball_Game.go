package baseballgame

import (
	"strconv"
)

func CalPoints(operations []string) int {
	num := []int{}
	i := 0
	for i < len(operations) {
		if operations[i] == "C" {
			operations = append(operations[:i], operations[i+1:]...)
			num = num[:len(num)-1]
		} else if operations[i] == "D" {
			operations = append(operations[:i], operations[i+1:]...)
			num = append(num, num[len(num)-1]*2)
		} else if operations[i] == "+" {
			operations = append(operations[:i], operations[i+1:]...)
			num = append(num, num[len(num)-1]+num[len(num)-2])
		} else {
			val, _ := strconv.Atoi(operations[i])
			num = append(num, val)
			i++
		}
	}
	sum := 0
	for index := 0; index < len(num); index++ {
		sum = sum + num[index]
	}
	return sum
}
