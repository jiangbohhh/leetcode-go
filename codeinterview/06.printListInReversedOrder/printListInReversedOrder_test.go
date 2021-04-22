package printListInReverseOrder

import (
	"jiangbohhh/leetcode-go/structure/linkedlist"
	"reflect"
	"testing"
)

func Test_reversePrint(t *testing.T) {
	type args struct {
		head []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "[1,3,2] => [2,3,1]",
			args: args{head: []int{1, 3, 2}},
			want: []int{2, 3, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			head := linkedlist.ListToNode(tt.args.head)
			if got := reversePrint(head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("reversePrint() = %v, want %v", got, tt.want)
			}
		})
	}
}
