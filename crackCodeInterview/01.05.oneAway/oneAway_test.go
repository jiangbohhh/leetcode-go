package oneAway

import "testing"

func Test_oneEditAway(t *testing.T) {
	type args struct {
		first  string
		second string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "pale,ple => true",
			args: args{
				first:  "pale",
				second: "ple",
			},
			want: true,
		},
		{
			name: "abcd,abc => true",
			args: args{
				first:  "abcd",
				second: "abc",
			},
			want: true,
		},
		{
			name: "abcd,bcd => true",
			args: args{
				first:  "abcd",
				second: "bcd",
			},
			want: true,
		},
		{
			name: "abcd,abed => true",
			args: args{
				first:  "abcd",
				second: "abed",
			},
			want: true,
		},
		{
			name: "mart,karma => false",
			args: args{
				first:  "mart",
				second: "karma",
			},
			want: false,
		},
		{
			name: "pales,pal => false",
			args: args{
				first:  "pales",
				second: "pal",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := oneEditAway(tt.args.first, tt.args.second); got != tt.want {
				t.Errorf("oneEditAway() = %v, want %v", got, tt.want)
			}
		})
	}
}
