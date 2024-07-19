package baseballgame

import "testing"

func TestCalPoints(t *testing.T) {
	type args struct {
		operations []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{name: "Test Case 1", args: args{operations: []string{"5", "2", "C", "D", "+"}}, want: 30},
		{name: "Test Case 2", args: args{operations: []string{"5", "-2", "4", "C", "D", "9", "+", "+"}}, want: 27},
		{name: "Test case 3", args: args{operations: []string{"1", "C"}}, want: 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CalPoints(tt.args.operations); got != tt.want {
				t.Errorf("CalPoints() = %v, want %v", got, tt.want)
			}
		})
	}
}
