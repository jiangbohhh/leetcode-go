package removeAllDuplicatesInString

import "testing"

func Test_removeDuplicates(t *testing.T) {
	type args struct {
		S string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "abbaca => ca",
			args: args{S: "abbaca"},
			want: "ca",
		},
		{
			name: "'' => ''",
			args: args{S: ""},
			want: "",
		},
		{
			name: "aaaaa => ''",
			args: args{S: "aaaaa"},
			want: "a",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeDuplicates(tt.args.S); got != tt.want {
				t.Errorf("removeDuplicates() = %v, want %v", got, tt.want)
			}
		})
	}
}
