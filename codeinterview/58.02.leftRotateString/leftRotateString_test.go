package leftRotateString

import "testing"

func Test_reverseLeftWords(t *testing.T) {
	type args struct {
		s string
		n int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: " '', 2 => ''",
			args: args{s: "", n: 2},
			want: "",
		},
		{
			name: "abcdefg, 0 => abcdefg",
			args: args{s: "abcdefg", n: 0},
			want: "abcdefg",
		},
		{
			name: "abcdefg, 1 => bcdefga",
			args: args{s: "abcdefg", n: 1},
			want: "bcdefga",
		},
		{
			name: "abcdefg, 2 => cdefgab",
			args: args{s: "abcdefg", n: 2},
			want: "cdefgab",
		},
		{
			name: "lrloseumgh, 6 => umghlrlose",
			args: args{s: "lrloseumgh", n: 6},
			want: "umghlrlose",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverseLeftWords(tt.args.s, tt.args.n); got != tt.want {
				t.Errorf("reverseLeftWords() = %v, want %v", got, tt.want)
			}
		})
	}
}
