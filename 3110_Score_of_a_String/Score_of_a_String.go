package scoreofastring

import "math"

func ScoreOfString(s string) int {
	//第一版
	// var result rune
	// var val2 rune
	// for i, val := range s {
	// 	if i > 0 {
	// 		if val > val2 {
	// 			result = result + val - val2
	// 		} else {
	// 			result = result + val2 - val
	// 		}
	// 		val2 = val
	// 	} else {
	// 		val2 = val
	// 	}
	// }
	// return int(result)

	//第二版
	// var result rune

	// for i := 0; i < len(s)-1; i++ {
	// 	result = result + abs(rune(s[i])-rune(s[i+1]))
	// }
	// return int(result)

	//第三版
	var result int
	for i := 0; i < len(s)-1; i++ {
		result = result + int(math.Abs(float64((rune(s[i]) - rune(s[i+1])))))
	}
	return result
}

//第二版func
// func abs(data rune) rune {
// 	if data < 0 {
// 		data = data * -1
// 	}
// 	return data
// }
