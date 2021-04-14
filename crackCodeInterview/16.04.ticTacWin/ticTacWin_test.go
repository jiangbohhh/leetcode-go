package ticTacWin

import "testing"

func Test_tictactoe(t *testing.T) {
	type args struct {
		board []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: " ['O X',' XO','X O'] => X",
			args: args{board: []string{"O X", " XO", "X O"}},
			want: "X",
		},
		{
			name: " ['OOX','XXO','OXO'] => Draw",
			args: args{board: []string{"OOX", "XXO", "OXO"}},
			want: "Draw",
		},
		{
			name: " ['OOX','XXO','OX '] => Pending",
			args: args{board: []string{"OOX", "XXO", "OX "}},
			want: "Pending",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tictactoe(tt.args.board); got != tt.want {
				t.Errorf("tictactoe() = %v, want %v", got, tt.want)
			}
		})
	}
}
