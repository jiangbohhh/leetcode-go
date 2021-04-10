package uglyNumber

import "testing"

func Test_isUgly(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "1 => true",
			args: args{n: 1},
			want: true,
		},
		{
			name: "-21 => false",
			args: args{n: -21},
			want: false,
		},
		{
			name: "0 => false",
			args: args{n: 0},
			want: false,
		},
		{
			name: "6 => true",
			args: args{n: 6},
			want: true,
		},
		{
			name: "8 => true",
			args: args{n: 8},
			want: true,
		},
		{
			name: "14 => false",
			args: args{n: 14},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isUgly(tt.args.n); got != tt.want {
				t.Errorf("isUgly() = %v, want %v", got, tt.want)
			}
		})
	}
}
