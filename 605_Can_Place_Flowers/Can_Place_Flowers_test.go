package canplaceflowers

import "testing"

func TestCanPlaceFlowers(t *testing.T) {
	type args struct {
		flowerbed []int
		n         int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{"test case 1", args{[]int{1, 0, 0, 0, 1}, 1}, true},
		{"test case 2", args{[]int{1, 0, 0, 0, 1}, 2}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CanPlaceFlowers(tt.args.flowerbed, tt.args.n); got != tt.want {
				t.Errorf("CanPlaceFlowers() = %v, want %v", got, tt.want)
			}
		})
	}
}
