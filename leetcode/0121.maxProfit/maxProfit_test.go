package maxProfit

import "testing"

func Test_maxProfit(t *testing.T) {
	type args struct {
		prices []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "[7,1,5,3,6,4] , 5",
			args: args{prices: []int{7, 1, 5, 3, 6, 4}},
			want: 5,
		},
		{
			name: "[1,2,3,2,1], 2",
			args: args{prices: []int{1,2,3,2,1}},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxProfit(tt.args.prices); got != tt.want {
				t.Errorf("maxProfit() = %v, want %v", got, tt.want)
			}
		})
	}
}
