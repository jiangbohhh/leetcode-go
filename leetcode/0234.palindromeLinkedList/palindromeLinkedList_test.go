package palindromeLinkedList

import (
	"jiangbohhh/leetcode-go/structure/linkedlist"
	"testing"
)

func Test_isPalindrome(t *testing.T) {
	type args struct {
		head []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "1->2",
			args: args{head: []int{1, 2}},
			want: false,
		},
		{
			name: "1->2->2->1",
			args: args{head: []int{1, 2, 2, 1}},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			head := linkedlist.ListToNode(tt.args.head)
			if got := isPalindrome(head); got != tt.want {
				t.Errorf("isPalindrome() = %v, want %v", got, tt.want)
			}
			head2 := linkedlist.ListToNode(tt.args.head)
			if got := isPalindrome2(head2); got != tt.want {
				t.Errorf("isPalindrome2() = %v, want %v", got, tt.want)
			}
		})
	}
}
