package removeDuplicates

import (
	"reflect"
	"testing"
)

func Test_removeDuplicates(t *testing.T) {
	type args struct {
		nums []int
	}
	type want struct {
		k    int
		nums []int
	}
	tests := []struct {
		name string
		args args
		want want
	}{
		{
			name: "[1,1,2] => 2, [1,2]",
			args: args{nums: []int{1, 1, 2}},
			want: want{
				k:    2,
				nums: []int{1, 2},
			},
		},
		{
			name: "[0,0,1,1,1,2,2,3,3,4] => 5, [0,1,2,3,4]",
			args: args{nums: []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}},
			want: want{
				k:    5,
				nums: []int{0, 1, 2, 3, 4},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeDuplicates(tt.args.nums); got != tt.want.k || !reflect.DeepEqual(tt.args.nums[:tt.want.k], tt.want.nums) {
				t.Errorf("removeDuplicates() = %v, want %v", tt.args.nums[:tt.want.k], tt.want)
			}
		})
	}
}
