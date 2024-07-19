package simplifypath

import "testing"

func TestSimplifyPath(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{name: "Test Case 1", args: args{path: "/home/"}, want: "/home"},
		{name: "Test Case 2", args: args{path: "/home//foo/"}, want: "/home/foo"},
		{name: "Test Case 3", args: args{path: "/home/user/Documents/../Pictures"}, want: "/home/user/Pictures"},
		{name: "Test Case 4", args: args{path: "/../"}, want: "/"},
		{name: "Test Case 5", args: args{path: "/.../a/../b/c/../d/./"}, want: "/.../b/d"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SimplifyPath(tt.args.path); got != tt.want {
				t.Errorf("SimplifyPath() = %v, want %v", got, tt.want)
			}
		})
	}
}
