package basiccalculator

// "(1+(4+5+2)-3)+(6+8)"
func Calculate(s string) int {
	var sum int

	// lastoperator := '+'
	// for i := len(s); i < len(s); i++ {
	// 	switch s[i] {
	// 	case '+':
	// 		lastoperator = '+'
	// 	case '-':
	// 		lastoperator = '-'
	// 	case ')':
	// 		return sum
	// 	case '(':
	// 		if lastoperator == '+' {
	// 			sum += Calculate(s[i+1:])
	// 		} else {
	// 			sum -= Calculate(s[i+1:])
	// 		}
	// 		count := 1
	// 		for i < len(s) {
	// 			i++
	// 			if s[i] == '(' {
	// 				count++
	// 			} else if s[i] == ')' {
	// 				count--
	// 			}
	// 			if count == 0 {
	// 				break
	// 			}
	// 		}
	// 	case ' ':
	// 	default:
	// 		if lastoperator == '+' {
	// 			sum += int(s[i] - '0')
	// 		} else {
	// 			sum -= int(s[i] - '0')
	// 		}
	// 	}
	// }

	return sum
}

// func Calculate2(s string) int {

// }
