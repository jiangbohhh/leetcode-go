package reorderArray

import (
	"reflect"
	"testing"
)

func Test_exchange(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "[1,2,3,4] => [1,3,2,4]",
			args: args{nums: []int{1, 2, 3, 4}},
			want: []int{1, 3, 2, 4},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := exchange(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("exchange() = %v, want %v", got, tt.want)
			}
		})
	}
}
