package mergestringsalternately

import "testing"

func TestMergeAlternately(t *testing.T) {
	type args struct {
		word1 string
		word2 string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MergeAlternately(tt.args.word1, tt.args.word2); got != tt.want {
				t.Errorf("MergeAlternately() = %v, want %v", got, tt.want)
			}
		})
	}
}
