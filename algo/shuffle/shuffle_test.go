package shuffle

import "testing"

func Test_shuffle(t *testing.T) {
	type args struct {
		list []int
		n    int
	}
	tests := []struct {
		name string
		args args
	}{
		{name: "[1,2,3,4,5,6,7], 7", args: args{
			list: []int{1, 2, 3, 4, 5, 6, 7},
			n:    7,
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			shuffle(tt.args.list, tt.args.n)
			t.Log(tt.args.list)
		})
	}
}
