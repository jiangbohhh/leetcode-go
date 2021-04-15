package addTwoNumber

import (
	"jiangbohhh/leetcode-go/structure/linkedlist"
	"reflect"
	"testing"
)

func Test_addTwoNumbers(t *testing.T) {
	type args struct {
		l1 []int
		l2 []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "[2,4,3],[5,6,4] => [7,0,8]",
			args: args{
				l1: []int{2, 4, 3},
				l2: []int{5, 6, 4},
			},
			want: []int{7, 0, 8},
		},
		{
			name: "[0],[0] => [0]",
			args: args{
				l1: []int{0},
				l2: []int{0},
			},
			want: []int{0},
		},
		{
			name: "[9,9,9,9,9,9,9], [9,9,9,9] => [8,9,9,9,0,0,0,1]",
			args: args{
				l1: []int{9, 9, 9, 9, 9, 9, 9},
				l2: []int{9, 9, 9, 9},
			},
			want: []int{8, 9, 9, 9, 0, 0, 0, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l1 := linkedlist.ListToNode(tt.args.l1)
			l2 := linkedlist.ListToNode(tt.args.l2)
			got := linkedlist.NodeToList(addTwoNumbers(l1, l2))
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("addTwoNumbers() = %v, want %v", got, tt.want)
			}
		})
	}
}
