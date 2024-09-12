package removingstarsfromastring

import "testing"

func TestRemoveStars(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{"Test case1", args{"leet**cod*e"}, "lecoe"},
		{"Test case2", args{"erase*****"}, ""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RemoveStars(tt.args.s); got != tt.want {
				t.Errorf("RemoveStars() = %v, want %v", got, tt.want)
			}
		})
	}
}
