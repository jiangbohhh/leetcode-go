package zeroMatrix

import (
	"reflect"
	"testing"
)

func Test_setZeros(t *testing.T) {
	type args struct {
		matrix [][]int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "[[1,1,1],[1,0,1],[1,1,1]] => [[1,0,1],[0,0,0],[1,0,1]]",
			args: args{
				matrix: [][]int{
					{1, 1, 1},
					{1, 0, 1},
					{1, 1, 1},
				},
			},
			want: [][]int{
				{1, 0, 1},
				{0, 0, 0},
				{1, 0, 1},
			},
		},
		{
			name: "[ [0,1,2,0],[3,4,5,2],[1,3,1,5]] => [[0,0,0,0],[0,4,5,0],[0,3,1,0]]",
			args: args{
				matrix: [][]int{
					{0, 1, 2, 0},
					{3, 4, 5, 2},
					{1, 3, 1, 5},
				},
			},
			want: [][]int{
				{0, 0, 0, 0},
				{0, 4, 5, 0},
				{0, 3, 1, 0},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			matrix := tt.args.matrix
			setZeros(matrix)
			if !reflect.DeepEqual(matrix, tt.want) {
				t.Errorf("setZeros() = %v want = %v", matrix, tt.want)
			}
		})
	}
}
