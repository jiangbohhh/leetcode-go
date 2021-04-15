package removeLinkedListElements

import (
	"reflect"
	"testing"
)

func Test_removeElements(t *testing.T) {
	type args struct {
		head []int
		val  int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "[1,2,6,3,4,5,6],6 => [1,2,3,4,5]",
			args: args{
				head: []int{1, 2, 6, 3, 4, 5, 6},
				val:  6,
			},
			want: []int{1, 2, 3, 4, 5},
		},
		{
			name: "[7,7,7,7],7 => []",
			args: args{
				head: []int{7, 7, 7, 7},
				val:  7,
			},
			want: []int{},
		},
		{
			name: "[],1 => []",
			args: args{
				head: []int{},
				val:  1,
			},
			want: []int{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var head *ListNode
			if len(tt.args.head) > 1 {
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
			got := removeElements(head, tt.args.val)

			gotList := make([]int, 0)
			for got != nil {
				gotList = append(gotList, got.Val)
				got = got.Next
			}
			if !reflect.DeepEqual(gotList, tt.want) {
				t.Errorf("removeElements() = %v, want %v", gotList, tt.want)
			}
		})
	}
}
