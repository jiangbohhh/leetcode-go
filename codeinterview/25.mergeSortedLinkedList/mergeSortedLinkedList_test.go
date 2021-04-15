package mergeSortedLinkedList

import (
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
		{
			name: "1->2->4, nil => 1->2->4",
			args: args{
				l1: []int{1, 2, 4},
				l2: []int{},
			},
			want: []int{1, 2, 4},
		},
		{
			name: "nil, 1->2->4 => 1->2->4",
			args: args{
				l1: []int{},
				l2: []int{1, 2, 4},
			},
			want: []int{1, 2, 4},
		},
		{
			name: "nil, nil => nil",
			args: args{
				l1: nil,
				l2: nil,
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l1 := listToNode(tt.args.l1)
			l2 := listToNode(tt.args.l2)
			head := mergeTwoLists(l1, l2)
			got := NodeToList(head)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mergeTwoLists() = %v, want %v", got, tt.want)
			}
		})
	}
}

// NodeToList
func NodeToList(head *ListNode) (list []int) {
	for head != nil {
		list = append(list, head.Val)
		head = head.Next
	}
	return
}

// listToNode
func listToNode(list []int) (head *ListNode) {
	if len(list) >= 1 {
		head = &ListNode{
			Val:  list[len(list)-1],
			Next: nil,
		}

	}
	for i := len(list) - 2; i >= 0; i-- {
		node := &ListNode{
			Val:  list[i],
			Next: nil,
		}
		node.Next = head
		head = node
	}
	return
}
