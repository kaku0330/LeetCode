package maximumnumberofwordsfoundinsentences

import "testing"

func TestMostWordsFound(t *testing.T) {
	type args struct {
		sentences []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{name: "Test Case 1", args: args{sentences: []string{"alice and bob love leetcode", "i think so too", "this is great thanks very much"}}, want: 6},
		{name: "Test Case 2", args: args{sentences: []string{"please wait", "continue to fight", "continue to win"}}, want: 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MostWordsFound(tt.args.sentences); got != tt.want {
				t.Errorf("MostWordsFound() = %v, want %v", got, tt.want)
			}
		})
	}
}
