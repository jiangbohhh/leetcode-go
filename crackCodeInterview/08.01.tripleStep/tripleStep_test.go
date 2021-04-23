package tripleStep

import "testing"

func Test_waysToStep(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "61->752119970",
			args: args{n: 61},
			want: 752119970,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := waysToStep(tt.args.n); got != tt.want {
				t.Errorf("waysToStep() = %v, want %v", got, tt.want)
			}
		})
	}
}
