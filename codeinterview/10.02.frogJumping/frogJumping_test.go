package frogJumping

import "testing"

func Test_numWays(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "0->1",
			args: args{n: 0},
			want: 1,
		},
		{
			name: "2->2",
			args: args{n: 2},
			want: 2,
		},
		{
			name: "7->21",
			args: args{n: 7},
			want: 21,
		},
		{
			name: "92->720754435",
			args: args{n: 92},
			want: 720754435,
		},
		{
			name: "95->93363621",
			args: args{n: 95},
			want: 93363621,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numWays(tt.args.n); got != tt.want {
				t.Errorf("numWays() = %v, want %v", got, tt.want)
			}
		})
	}
}
