package scoreofastring

import "testing"

func TestScoreOfString(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{name: "Test Case 1", args: args{s: "hello"}, want: 13},
		{name: "Test Case 2", args: args{s: "zaz"}, want: 50},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ScoreOfString(tt.args.s); got != tt.want {
				t.Errorf("ScoreOfString() = %v, want %v", got, tt.want)
			}
		})
	}
}
