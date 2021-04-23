package calculator

import "testing"

func Test_calculate(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "3+2*2 => 7",
			args: args{s: "3+2*2"},
			want: 7,
		},
		{
			name: "3+20*2 => 43",
			args: args{s: "3+20*2"},
			want: 43,
		},
		{
			name: "3/2 => 1",
			args: args{s: "3/2"},
			want: 1,
		},
		{
			name: "0-2147483647 => -2147483647",
			args: args{s: "0-2147483647"},
			want: -2147483647,
		},
		{
			name: "1-1+1 => 1",
			args: args{s: "1-1+1"},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := calculate(tt.args.s); got != tt.want {
				t.Errorf("calculate() = %v, want %v", got, tt.want)
			}
		})
	}
}
