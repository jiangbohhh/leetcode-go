package climbingStairs

import "testing"

func Test_climbStairs(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "1 => 1",
			args: args{1},
			want: 1,
		},
		{
			name: "2 => 2",
			args: args{2},
			want: 2,
		},
		{
			name: "3 => 3",
			args: args{3},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := climbStairs(tt.args.n); got != tt.want {
				t.Errorf("climbStairs() = %v, want %v", got, tt.want)
			}
		})
	}
}
