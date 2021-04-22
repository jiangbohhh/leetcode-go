package reverseList

import (
	"jiangbohhh/leetcode-go/structure/linkedlist"
	"reflect"
	"testing"
)

func Test_reverseList(t *testing.T) {
	type args struct {
		head []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "1->2->3->4->5 => 5->4->3->2->1",
			args: args{head: []int{1, 2, 3, 4, 5}},
			want: []int{5, 4, 3, 2, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			head := linkedlist.ListToNode(tt.args.head)
			if got := linkedlist.NodeToList(reverseList(head)); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("reverseList() = %v, want %v", got, tt.want)
			}
		})
	}
}
