package findgreatestcommondivisorofarray

import "testing"

func TestFindGCD(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{"test case 1", args{[]int{2, 5, 6, 9, 10}}, 2},
		{"test case 2", args{[]int{7, 5, 6, 8, 3}}, 1},
		{"test case 3", args{[]int{3, 3}}, 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FindGCD(tt.args.nums); got != tt.want {
				t.Errorf("FindGCD() = %v, want %v", got, tt.want)
			}
		})
	}
}
