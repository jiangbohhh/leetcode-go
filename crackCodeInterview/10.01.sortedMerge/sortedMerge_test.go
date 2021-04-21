package sortedMerge

import (
	"reflect"
	"testing"
)

func Test_merge(t *testing.T) {
	type args struct {
		A []int
		m int
		B []int
		n int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "[1,2,3,0,0,0], [2,5,6] => [1,2,2.3,5,6]",
			args: args{
				A: []int{1, 2, 3, 0, 0, 0},
				m: 3,
				B: []int{2, 5, 6},
				n: 3,
			},
			want: []int{1, 2, 2, 3, 5, 6},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			merge(tt.args.A, tt.args.m, tt.args.B, tt.args.n)
			if !reflect.DeepEqual(tt.args.A, tt.want) {
				t.Errorf("merge() = %v, want = %v", tt.args.A, tt.want)
			}
		})
	}
}
