package permutationdifferencebetweentwostrings

import "testing"

func TestFindPermutationDifference(t *testing.T) {
	type args struct {
		s string
		t string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{name: "Test Case 1", args: args{s: "abc", t: "bac"}, want: 2},
		{name: "Test Case 2", args: args{s: "abcde", t: "edbac"}, want: 12},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FindPermutationDifference(tt.args.s, tt.args.t); got != tt.want {
				t.Errorf("FindPermutationDifference() = %v, want %v", got, tt.want)
			}
		})
	}
}
