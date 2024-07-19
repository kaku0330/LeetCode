package firstuniquecharacterinastring

import "testing"

func TestFirstUniqChar(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{name: "Test Case 1", args: args{s: "leetcode"}, want: 0},
		{name: "Test Case 2", args: args{s: "loveleetcode"}, want: 2},
		{name: "Test Case 3", args: args{s: "aabb"}, want: -1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FirstUniqChar(tt.args.s); got != tt.want {
				t.Errorf("FirstUniqChar() = %v, want %v", got, tt.want)
			}
		})
	}
}
