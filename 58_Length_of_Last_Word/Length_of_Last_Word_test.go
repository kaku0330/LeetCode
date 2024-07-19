package lengthoflastword

import "testing"

func TestLengthOfLastWord(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{name: "Test Case 1", args: args{s: "Hello World"}, want: 5},
		{name: "Test Case 2", args: args{s: "   fly me   to   the moon  "}, want: 4},
		{name: "Test Case 3", args: args{s: "luffy is still joyboy"}, want: 6},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LengthOfLastWord(tt.args.s); got != tt.want {
				t.Errorf("LengthOfLastWord() = %v, want %v", got, tt.want)
			}
		})
	}
}
