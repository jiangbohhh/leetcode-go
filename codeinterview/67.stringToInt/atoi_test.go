package stringToInt

import "testing"

func Test_strToInt(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "'42' => 42",
			args: args{"42"},
			want: 42,
		},
		{
			name: "'  -42' => -42",
			args: args{"   -42"},
			want: -42,
		},
		{
			name: "'4193 with words' => 4193",
			args: args{"4193 with words"},
			want: 4193,
		},
		{
			name: "'words and 987' => 0",
			args: args{"words and 987"},
			want: 0,
		},
		{
			name: "'-91283472332' => -2147483648",
			args: args{"-91283472332"},
			want: -2147483648,
		},
		{
			name: "'9223372036854775808' => 2147483647",
			args: args{"9223372036854775808"},
			want: 2147483647,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := strToInt(tt.args.str); got != tt.want {
				t.Errorf("strToInt() = %v, want %v", got, tt.want)
			}
		})
	}
}
