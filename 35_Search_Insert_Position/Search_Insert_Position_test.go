package searchinsertposition

import "testing"

func TestSearchInsert(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{name: "Test Case 1", args: args{nums: []int{1, 3, 5, 6}, target: 5}, want: 2},
		{name: "Test Case 2", args: args{nums: []int{1, 3, 5, 6}, target: 2}, want: 1},
		{name: "Test Case 3", args: args{nums: []int{1, 3, 5, 6}, target: 7}, want: 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SearchInsert(tt.args.nums, tt.args.target); got != tt.want {
				t.Errorf("SearchInsert() = %v, want %v", got, tt.want)
			}
		})
	}
}
