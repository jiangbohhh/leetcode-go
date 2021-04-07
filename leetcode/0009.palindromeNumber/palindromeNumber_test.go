package palindromeNumber

import "testing"

func Test_isPalindrome(t *testing.T) {
	type args struct {
		x int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "121=>true",
			args: args{121},
			want: true,
		},
		{
			name: "-121=>false",
			args: args{-121},
			want: false,
		},
		{
			name: "10=>false",
			args: args{10},
			want: false,
		},
		{
			name: "-101=>false",
			args: args{-101},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPalindrome(tt.args.x); got != tt.want {
				t.Errorf("isPalindrome() = %v, want %v", got, tt.want)
			}
			if got := isPalindromeNum(tt.args.x); got != tt.want {
				t.Errorf("isPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}
