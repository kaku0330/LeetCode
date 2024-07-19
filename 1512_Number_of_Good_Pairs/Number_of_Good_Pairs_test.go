package numberofgoodpairs

import "testing"

func TestNumIdenticalPairs(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{name: "Test Case 1", args: args{nums: []int{1, 2, 3, 1, 1, 3}}, want: 4},
		{name: "Test Case 2", args: args{nums: []int{1, 1, 1, 1}}, want: 6},
		{name: "Test Case 3", args: args{nums: []int{1, 2, 3}}, want: 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NumIdenticalPairs(tt.args.nums); got != tt.want {
				t.Errorf("NumIdenticalPairs() = %v, want %v", got, tt.want)
			}
		})
	}
}
