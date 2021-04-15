package kthNodeFromEnd

import (
	"jiangbohhh/leetcode-go/structure/linkedlist"
	"reflect"
	"testing"
)

func Test_getKthFromEnd(t *testing.T) {
	type args struct {
		head []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "1->2->3->4->5, k=2 => 4->5",
			args: args{
				head: []int{1, 2, 3, 4, 5},
				k:    2,
			},
			want: []int{4, 5},
		},
		{
			name: "nil, k = 2 => nil",
			args: args{
				head: []int{},
				k:    2,
			},
			want: nil,
		},
		{
			name: "1->2->3->4->5, k=6 => 1->2->3->4->5",
			args: args{
				head: []int{1, 2, 3, 4, 5},
				k:    6,
			},
			want: []int{1, 2, 3, 4, 5},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			head := linkedlist.ListToNode(tt.args.head)
			got := linkedlist.NodeToList(getKthFromEnd(head, tt.args.k))
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getKthFromEnd() = %v, want %v", got, tt.want)
			}
		})
	}
}
