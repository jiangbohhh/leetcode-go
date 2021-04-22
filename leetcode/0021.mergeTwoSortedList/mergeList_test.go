package mergeTwoSortedList

import (
	"jiangbohhh/leetcode-go/structure/linkedlist"
	"reflect"
	"testing"
)

func Test_mergeTwoLists(t *testing.T) {
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
			name: "1->2->4, 1->3->4 => 1->1->2->3->4->4",
			args: args{
				l1: []int{1, 2, 4},
				l2: []int{1, 3, 4},
			},
			want: []int{1, 1, 2, 3, 4, 4},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l1 := linkedlist.ListToNode(tt.args.l1)
			l2 := linkedlist.ListToNode(tt.args.l2)
			if got := linkedlist.NodeToList(mergeTwoLists(l1, l2)); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mergeTwoLists() = %v, want %v", got, tt.want)
			}
		})
	}
}
