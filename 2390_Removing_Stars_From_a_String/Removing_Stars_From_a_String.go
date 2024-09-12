package removingstarsfromastring

func RemoveStars(s string) string {
	// spliteS := strings.Split(s, "")
	// for i := 0; i < len(spliteS); {
	// 	if spliteS[i] == "*" {
	// 		spliteS = append(spliteS[:i-1], spliteS[i:]...)
	// 		spliteS = append(spliteS[:i-1], spliteS[i:]...)
	// 		i--
	// 	} else {
	// 		i++
	// 	}
	// }
	// return strings.Join(spliteS, "")

	result := []byte{}
	for i := 0; i < len(s); i++ {
		if s[i] == 42 {
			result = result[:len(result)-1]
		} else {
			result = append(result, s[i])
		}
	}

	return string(result)
}
