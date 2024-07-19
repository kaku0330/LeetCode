package jewelsandstones

import "testing"

func TestNumJewelsInStones(t *testing.T) {
	type args struct {
		jewels string
		stones string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{name: "Test case 1", args: args{jewels: "aA", stones: "aAAbbbb"}, want: 3},
		{name: "Test case 2", args: args{jewels: "z", stones: "ZZ"}, want: 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NumJewelsInStones(tt.args.jewels, tt.args.stones); got != tt.want {
				t.Errorf("NumJewelsInStones() = %v, want %v", got, tt.want)
			}
		})
	}
}
