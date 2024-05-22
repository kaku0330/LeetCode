package numberofstudentsunabletoeatlunch

//students, sandwiches := []int{1, 1, 0, 0}, []int{0, 1, 0, 1}
func CountStudents(students []int, sandwiches []int) int {
	// 第一版
	// count := 0
	// for count < len(students) {
	// 	if students[0] == sandwiches[0] {
	// 		sandwiches = sandwiches[1:]
	// 		students = students[1:]
	// 		count = 0
	// 	} else {
	// 		students = append(students, students[0])
	// 		students = students[1:]
	// 		count++
	// 	}
	// 	fmt.Println(students)
	// 	fmt.Println(sandwiches)
	// }
	// return len(students)

	return len(students)
}
