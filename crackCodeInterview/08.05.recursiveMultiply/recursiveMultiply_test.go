package recursiveMultiply

import "testing"

func Test_multiply(t *testing.T) {
	type args struct {
		A int
		B int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "1, 10 => 10",
			args: args{
				A: 1,
				B: 10,
			},
			want: 10,
		},
		{
			name: "3, 4 => 12",
			args: args{
				A: 3,
				B: 4,
			},
			want: 12,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := multiply(tt.args.A, tt.args.B); got != tt.want {
				t.Errorf("multiply() = %v, want %v", got, tt.want)
			}
		})
	}
}
