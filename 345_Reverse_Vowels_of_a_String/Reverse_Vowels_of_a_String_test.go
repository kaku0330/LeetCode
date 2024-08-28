package reversevowelsofastring

import "testing"

func TestReverseVowels(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{name: "test case 1", args: args{s: "hello"}, want: "holle"},
		{name: "test case 2", args: args{s: "leetcode"}, want: "leotcede"},
		{name: "test case 3", args: args{s: "aA"}, want: "Aa"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ReverseVowels(tt.args.s); got != tt.want {
				t.Errorf("ReverseVowels() = %v, want %v", got, tt.want)
			}
		})
	}
}
