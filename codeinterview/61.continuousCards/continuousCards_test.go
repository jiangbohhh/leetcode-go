package continuousCards

import "testing"

func Test_isStraight(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "[1,2,3,4,5] => true",
			args: args{nums: []int{1, 2, 3, 4, 5}},
			want: true,
		},
		{
			name: "[0,0,1,2,5] => true",
			args: args{nums: []int{0, 0, 1, 2, 5}},
			want: true,
		},
		{
			name: "[1,2,5,7,8] => false",
			args: args{nums: []int{1, 2, 5, 7, 8}},
			want: false,
		},
		{
			name: "[1,10,11,12,13] => true",
			args: args{nums: []int{1, 10, 11, 12, 13}},
			want: false,
		},
		{
			name: "[0,0,2,2,5] => false",
			args: args{nums: []int{0, 0, 2, 2, 5}},
			want: false,
		},
		{
			name: "[10,11,0,12,6] => false",
			args: args{nums: []int{10, 11, 0, 12, 6}},
			want: false,
		},

		{
			name: "[12,10,1,5,13] => false",
			args: args{nums: []int{12, 10, 1, 5, 13}},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isStraight(tt.args.nums); got != tt.want {
				t.Errorf("isStraight() = %v, want %v", got, tt.want)
			}
		})
	}
}
