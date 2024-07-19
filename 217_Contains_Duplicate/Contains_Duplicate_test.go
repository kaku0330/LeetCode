package containsduplicate

import "testing"

func TestContainsDuplicate(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{name: "Test Case 1", args: args{nums: []int{1, 2, 3, 1}}, want: true},
		{name: "Test Case 2", args: args{nums: []int{1, 2, 3, 4}}, want: false},
		{name: "Test Case 3", args: args{nums: []int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2}}, want: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ContainsDuplicate(tt.args.nums); got != tt.want {
				t.Errorf("ContainsDuplicate() = %v, want %v", got, tt.want)
			}
		})
	}
}
