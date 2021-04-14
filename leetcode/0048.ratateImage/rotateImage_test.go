package ratateImage

import (
	"reflect"
	"testing"
)

func Test_rotate(t *testing.T) {
	type args struct {
		matrix [][]int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "[[1,2,3],[4,5,6],[7,8,9]] => [[7,4,1],[8,5,2],[9,6,3]]",
			args: args{
				matrix: [][]int{
					{1, 2, 3},
					{4, 5, 6},
					{7, 8, 9},
				},
			},
			want: [][]int{
				{7, 4, 1},
				{8, 5, 2},
				{9, 6, 3},
			},
		},
		{
			name: "[[1,2],[3,4]] => [[3,1],[4,2]]",
			args: args{
				matrix: [][]int{
					{1, 2},
					{3, 4},
				},
			},
			want: [][]int{
				{3, 1},
				{4, 2},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			matrix := tt.args.matrix
			rotate(matrix)
			if !reflect.DeepEqual(matrix, tt.want) {
				t.Errorf("rotate() = %v want = %v", matrix, tt.want)
			}
		})
	}
}
