package decodeWays

import "testing"

func Test_numDecodings(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "12=>2",
			args: args{s: "12"},
			want: 2,
		},
		{
			name: "226=>3",
			args: args{s: "226"},
			want: 3,
		},
		{
			name: "0=>0",
			args: args{s: "0"},
			want: 0,
		},
		{
			name: "12258=5",
			args: args{s: "12258"},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numDecodings(tt.args.s); got != tt.want {
				t.Errorf("numDecodings() = %v, want %v", got, tt.want)
			}
		})
	}
}
