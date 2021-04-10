package masterMind

import (
	"reflect"
	"testing"
)

func Test_masterMind(t *testing.T) {
	type args struct {
		solution string
		guess    string
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "RRRR,GGGG => [0,0]",
			args: args{
				solution: "RRRR",
				guess:    "GGGG",
			},
			want: []int{0, 0},
		},
		{
			name: "RRRR,RRRR => [4,0]",
			args: args{
				solution: "RRRR",
				guess:    "RRRR",
			},
			want: []int{4, 0},
		},
		{
			name: "RYGB,BGYR => [0,4]",
			args: args{
				solution: "RYGB",
				guess:    "BGYR",
			},
			want: []int{0, 4},
		},
		{
			name: "RGBY,GGRR => [1,1]",
			args: args{
				solution: "RGBY",
				guess:    "GGRR",
			},
			want: []int{1, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := masterMind(tt.args.solution, tt.args.guess); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("masterMind() = %v, want %v", got, tt.want)
			}
		})
	}
}
