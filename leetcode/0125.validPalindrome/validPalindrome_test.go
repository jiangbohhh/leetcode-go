package validPalindrome

import "testing"

func Test_isPalindrome(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: " => true",
			args: args{s: ""},
			want: true,
		},
		{
			name: "   => true",
			args: args{s: "  "},
			want: true,
		},
		{
			name: "A man, a plan, a canal: Panama => true",
			args: args{s: "A man, a plan, a canal: Panama"},
			want: true,
		},
		{
			name: "race a car => false",
			args: args{s: "race a car"},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPalindrome(tt.args.s); got != tt.want {
				t.Errorf("isPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}
