package minstack

import (
	"reflect"
	"testing"
)

func TestConstructor(t *testing.T) {
	tests := []struct {
		name string
		want MinStack
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Constructor(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Constructor() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMinStack_Push(t *testing.T) {
	type args struct {
		val int
	}
	tests := []struct {
		name string
		this *MinStack
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.this.Push(tt.args.val)
		})
	}
}

func TestMinStack_Pop(t *testing.T) {
	tests := []struct {
		name string
		this *MinStack
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.this.Pop()
		})
	}
}

func TestMinStack_Top(t *testing.T) {
	tests := []struct {
		name string
		this *MinStack
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.this.Top(); got != tt.want {
				t.Errorf("MinStack.Top() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMinStack_GetMin(t *testing.T) {
	tests := []struct {
		name string
		this *MinStack
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.this.GetMin(); got != tt.want {
				t.Errorf("MinStack.GetMin() = %v, want %v", got, tt.want)
			}
		})
	}
}
