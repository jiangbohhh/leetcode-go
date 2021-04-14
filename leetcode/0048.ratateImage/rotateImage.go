package ratateImage

// 0048. Rotate Image
// 给定一个 n×n 的二维矩阵matrix 表示一个图像。请你将图像顺时针旋转 90 度。
// 你必须在 原地 旋转图像，这意味着你需要直接修改输入的二维矩阵。请不要 使用另一个矩阵来旋转图像。
//
// 示例 1：
// 输入：matrix = [[1,2,3],[4,5,6],[7,8,9]]
// 输出：[[7,4,1],[8,5,2],[9,6,3]]
// 示例 2：
// 输入：matrix = [[5,1,9,11],[2,4,8,10],[13,3,6,7],[15,14,12,16]]
// 输出：[[15,13,2,5],[14,3,4,1],[12,6,8,9],[16,7,10,11]]
// 示例 3：
// 输入：matrix = [[1]]
// 输出：[[1]]
// 示例 4：
// 输入：matrix = [[1,2],[3,4]]
// 输出：[[3,1],[4,2]]
//
// 提示：
// matrix.length == n
// matrix[i].length == n
// 1 <= n <= 20
// -1000 <= matrix[i][j] <= 1000
//

// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/rotate-image
// 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

// 一般这种题目有多种解法：1. 借助辅助矩阵。2. 翻转代替旋转。 3. 原地交换

// rotate
func rotate(matrix [][]int) {
	n := len(matrix)
	// 旋转 90° 相当于上下翻转(180°)后接左上右下对角线翻转(270°)
	for i := 0; i < n>>1; i++ {
		for j := 0; j < n; j++ {
			matrix[i][j], matrix[n-1-i][j] = matrix[n-1-i][j], matrix[i][j]
		}
	}
	// 左上右下对角线翻转
	for i := 0; i < n; i++ {
		for j := 0; j < i; j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}
}

// (0,0) -> (0, n-1)
//     |       |
// (n-1, 0)  (n-1, n-1)
// 交替替换外圈的值

// rotate2
func rotate2(matrix [][]int) {
	// 原地交换做法解决
	n := len(matrix)
	i1, j1 := 0, 0
	for n > 1 {
		i2, j2 := i1, j1+n-1
		i3, j3 := i1+n-1, j1+n-1
		i4, j4 := i1+n-1, j1
		for move := 0; move <= n-2; move++ {
			// 从左到右 所以 j1+即可
			jj1 := j1 + move
			// 从上到下 i2+即可
			ii2 := i2 + move
			// 从右到左 j3-即可
			jj3 := j3 - move
			// 从下到上 i4—即可
			ii4 := i4 - move
			matrix[i1][jj1], matrix[ii2][j2], matrix[i3][jj3], matrix[ii4][j4] =
				matrix[ii4][j4], matrix[i1][jj1], matrix[ii2][j2], matrix[i3][jj3]
		}
		i1++
		j1++
		n -= 2
	}
}
