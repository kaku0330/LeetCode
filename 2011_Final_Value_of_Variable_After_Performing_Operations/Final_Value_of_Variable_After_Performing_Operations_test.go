package finalvalueofvariableafterperformingoperations

import "testing"

func TestFinalValueAfterOperations(t *testing.T) {
	type args struct {
		operations []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{name: "Test Case 1", args: args{operations: []string{"--X", "X++", "X++"}}, want: 1},
		{name: "Test Case 2", args: args{operations: []string{"++X", "++X", "X++"}}, want: 3},
		{name: "Test Case 3", args: args{operations: []string{"X++", "++X", "--X", "X--"}}, want: 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FinalValueAfterOperations(tt.args.operations); got != tt.want {
				t.Errorf("FinalValueAfterOperations() = %v, want %v", got, tt.want)
			}
		})
	}
}
