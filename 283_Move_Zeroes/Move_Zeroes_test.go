package movezeroes

import (
	"reflect"
	"testing"
)

func TestMoveZeroes(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
		{"Test case1", args{[]int{0, 1, 0, 3, 12}}, []int{1, 3, 12, 0, 0}},
		{"Test case2", args{[]int{0, 0, 1}}, []int{1, 0, 0}},
		{"Test case3", args{[]int{0, 0, 0}}, []int{0, 0, 0}},
		{"Test case4", args{[]int{4, 2, 4, 0, 0, 3, 0, 5, 1, 0}}, []int{4, 2, 4, 3, 5, 1, 0, 0, 0, 0}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MoveZeroes(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MoveZeroes() = %v, want %v", got, tt.want)
			}
		})
	}
}
