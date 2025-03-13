package partitioningintominimumnumberofdecibinarynumbers

import "testing"

func MinPartitions_Test(t *testing.T) {
	type args struct {
		n string
	}

	tests := []struct {
		msg  string
		args args
		want int
	}{
		{msg: "Test Case 1", args: args{n: "32"}, want: 3},
		{msg: "Test Case 2", args: args{n: "82734"}, want: 8},
		{msg: "Test Case 3", args: args{n: "27346209830709182346"}, want: 9},
	}

	for _, testdata := range tests {
		t.Run(testdata.msg, func(t *testing.T) {
			if got := MinPartitions(testdata.args.n); got != testdata.want {
				t.Errorf("MinPartitions() = %v, want %v", got, testdata.want)
			}
		})

	}
}
