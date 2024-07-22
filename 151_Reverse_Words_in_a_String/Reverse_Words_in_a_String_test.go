package reversewordsinastring

import "testing"

func TestReverseWords(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{name: "Test Case 1", args: args{s: "the sky is blue"}, want: "blue is sky the"},
		{name: "Test Case 2", args: args{s: "  hello world  "}, want: "world hello"},
		{name: "Test Case 3", args: args{s: "a good   example"}, want: "example good a"},
		{name: "Test Case 4", args: args{s: " 3c      2zPeO dpIMVv2SG    1AM       o       VnUhxK a5YKNyuG     x9    EQ  ruJO       0Dtb8qG91w 1rT3zH F0m n G wU"}, want: "wU G n F0m 1rT3zH 0Dtb8qG91w ruJO EQ x9 a5YKNyuG VnUhxK o 1AM dpIMVv2SG 2zPeO 3c"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ReverseWords(tt.args.s); got != tt.want {
				t.Errorf("ReverseWords() = %v, want %v", got, tt.want)
			}
		})
	}
}

//wU G n F0m 1rT3zH 0Dtb8qG91w ruJO EQ x9 a5YKNyuG VnUhxK o 1AM dpIMVv2SG 2zPeO 3c
//wU G n F0m 1rT3zH 0Dtb8qG91w ruJO EQ  x9 a5YKNyuG VnUhxK o 1AM dpIMVv2SG 2zPeO 3c
