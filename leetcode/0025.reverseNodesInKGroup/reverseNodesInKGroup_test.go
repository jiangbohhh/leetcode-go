package reverseNodesInKGroup

import (
	"jiangbohhh/leetcode-go/structure/linkedlist"
	"reflect"
	"testing"
)

func Test_reverseKGroup(t *testing.T) {
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
			name: "[1,2,3,4,5] k=2 => [2,1,4,3,5]",
			args: args{
				head: []int{1, 2, 3, 4, 5},
				k:    2,
			},
			want: []int{2, 1, 4, 3, 5},
		},
		{
			name: "[1,2,3,4,5] k=3 => [3,2,1,4,5]",
			args: args{
				head: []int{1, 2, 3, 4, 5},
				k:    3,
			},
			want: []int{3, 2, 1, 4, 5},
		},
		{
			name: "[1,2,3,4,5] k=1 => [1,2,3,4,5]",
			args: args{
				head: []int{1, 2, 3, 4, 5},
				k:    1,
			},
			want: []int{1, 2, 3, 4, 5},
		},
		{
			name: "[1] k=1 => [1]",
			args: args{
				head: []int{1},
				k:    1,
			},
			want: []int{1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			head := linkedlist.ListToNode(tt.args.head)
			if got := linkedlist.NodeToList(reverseKGroup(head, tt.args.k)); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("reverseKGroup() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_reverseKGroup1(t *testing.T) {
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
			name: "[1,2,3,4,5] k=2 => [2,1,4,3,5]",
			args: args{
				head: []int{1, 2, 3, 4, 5},
				k:    2,
			},
			want: []int{2, 1, 4, 3, 5},
		},
		{
			name: "[1,2,3,4,5] k=3 => [3,2,1,4,5]",
			args: args{
				head: []int{1, 2, 3, 4, 5},
				k:    3,
			},
			want: []int{3, 2, 1, 4, 5},
		},
		{
			name: "[1,2,3,4,5] k=1 => [1,2,3,4,5]",
			args: args{
				head: []int{1, 2, 3, 4, 5},
				k:    1,
			},
			want: []int{1, 2, 3, 4, 5},
		},
		{
			name: "[1] k=1 => [1]",
			args: args{
				head: []int{1},
				k:    1,
			},
			want: []int{1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			head := linkedlist.ListToNode(tt.args.head)
			if got := linkedlist.NodeToList(reverseKGroup1(head, tt.args.k)); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("reverseKGroup() = %v, want %v", got, tt.want)
			}
		})
	}
}
