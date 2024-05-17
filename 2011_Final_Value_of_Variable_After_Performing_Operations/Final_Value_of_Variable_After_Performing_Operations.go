package finalvalueofvariableafterperformingoperations

func FinalValueAfterOperations(operations []string) int {
	result := 0
	for i := 0; i < len(operations); i++ {
		if operations[i] == "--X" {
			result -= 1
		} else if operations[i] == "X--" {
			result -= 1
		} else if operations[i] == "++X" {
			result += 1
		} else if operations[i] == "X++" {
			result += 1
		}
	}

	return result
}
