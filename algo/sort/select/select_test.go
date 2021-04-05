package _select

import (
	"reflect"
	"testing"
)

func Test_selectSort(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "[4,5,6,1,2,3] => [1,2,3,4,5,6]",
			args: args{nums: []int{4,5,6,1,2,3}},
			want: []int{1,2,3,4,5,6},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			nums := tt.args.nums
			selectSort(nums)
			if !reflect.DeepEqual(nums,tt.want) {
				t.Errorf("selectSort = %v, want = %v", nums, tt.want)
			}
		})
	}
}
