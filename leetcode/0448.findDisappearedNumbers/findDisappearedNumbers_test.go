package findDisappearedNumbers

import (
	"reflect"
	"testing"
)

func Test_findDisappearedNumbers(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "[4,3,2,7,8,2,3,1], [5,6]",
			args: args{nums: []int{4, 3, 2, 7, 8, 2, 3, 1}},
			want: []int{5, 6},
		},
		{
			name: "[4,3,2,5,6,2,3,1], [7,8]",
			args: args{nums: []int{4, 3, 2, 5, 6, 2, 3, 1}},
			want: []int{7, 8},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findDisappearedNumbers(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findDisappearedNumbers() = %v, want %v", got, tt.want)
			}
		})
	}
}
