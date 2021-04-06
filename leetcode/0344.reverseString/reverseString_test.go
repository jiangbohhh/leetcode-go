package reverseString

import "testing"

func Test_reverseString(t *testing.T) {
	type args struct {
		s []byte
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "hello => olleh",
			args: args{s: []byte{'h', 'e', 'l', 'l', 'o'}},
			want: "olleh",
		},
		{
			name: "Hannah => hannaH",
			args: args{s: []byte{'H', 'a', 'n', 'n', 'a', 'h'}},
			want: "hannaH",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := tt.args.s
			reverseString(s)
			if string(s) != tt.want {
				t.Errorf("reverseString() = %v, want = %v", s, tt.want)
			}
		})
	}
}
