package middleOfTheLinkedList

import (
	"reflect"
	"testing"
)

func Test_middleNode(t *testing.T) {
	type args struct {
		head []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "[1,2,3,4,5] => [3,4,5]",
			args: args{head: []int{1, 2, 3, 4, 5}},
			want: []int{3, 4, 5},
		},
		{
			name: "[1,2,3,4,5,6] => [4,5,6]",
			args: args{head: []int{1, 2, 3, 4, 5, 6}},
			want: []int{4, 5, 6},
		},
		{
			name: "[1,] => [1]",
			args: args{head: []int{1}},
			want: []int{1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var head *ListNode
			if len(tt.args.head) >= 1 {
				head = &ListNode{
					Val:  tt.args.head[len(tt.args.head)-1],
					Next: nil,
				}

			}
			for i := len(tt.args.head) - 2; i >= 0; i-- {
				node := &ListNode{
					Val:  tt.args.head[i],
					Next: nil,
				}
				node.Next = head
				head = node
			}
			got := middleNode(head)
			gotList := make([]int, 0)
			for got != nil {
				gotList = append(gotList, got.Val)
				got = got.Next
			}
			if !reflect.DeepEqual(gotList, tt.want) {
				t.Errorf("middleNode() = %v, want %v", gotList, tt.want)
			}
		})
	}
}
