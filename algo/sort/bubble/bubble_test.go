package bubble

import (
	"jiangbohhh/leetcode-go/algo/sort"
	"reflect"
	"testing"
)

func Test_bubbleSort(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "[4,5,6,1,2,3] => [1,2,3,4,5,6]",
			args: args{nums: []int{4, 5, 6, 1, 2, 3}},
			want: []int{1, 2, 3, 4, 5, 6},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			nums := tt.args.nums
			bubbleSort(nums)
			if !reflect.DeepEqual(nums, tt.want) {
				t.Errorf("bubbleSort() = %v, want = %v \n", nums, tt.want)
			}
		})
	}
}

func Benchmark_bubbleSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sequence := sort.RandomSequence(10000)
		bubbleSort(sequence)
	}
}
