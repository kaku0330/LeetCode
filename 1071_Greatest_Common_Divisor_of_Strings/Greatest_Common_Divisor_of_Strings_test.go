package greatestcommondivisorofstrings

import "testing"

func TestGcdOfStrings(t *testing.T) {
	type args struct {
		str1 string
		str2 string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "Test Case 1", args: args{str1: "ABCABC", str2: "ABC"}, want: "ABC"},
		{name: "Test Case 2", args: args{str1: "ABABAB", str2: "ABAB"}, want: "AB"},
		{name: "Test Case 3", args: args{str1: "LEET", str2: "CODE"}, want: ""},
		{name: "Test Case 4", args: args{str1: "ABCDEF", str2: "ABC"}, want: ""},
		{name: "Test Case 5", args: args{str1: "TAUXXTAUXXTAUXXTAUXXTAUXX", str2: "TAUXXTAUXXTAUXXTAUXXTAUXXTAUXXTAUXXTAUXXTAUXX"}, want: "TAUXX"},
		{name: "Test Case 6", args: args{str1: "ABABABAB", str2: "ABAB"}, want: "ABAB"},
		{name: "Test Case 7", args: args{str1: "AA", str2: "A"}, want: "A"},
		{name: "Test Case 8", args: args{str1: "AAAAAAAAA", str2: "AAACCC"}, want: ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GcdOfStrings(tt.args.str1, tt.args.str2); got != tt.want {
				t.Errorf("GcdOfStrings() = %v, want %v", got, tt.want)
			}
		})
	}
}
