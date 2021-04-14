package spiralMatrix

// 0054. 螺旋矩阵
// 给你一个 m 行 n 列的矩阵 matrix ，请按照 顺时针螺旋顺序 ，返回矩阵中的所有元素。
// 示例 1：
// 输入：matrix = [[1,2,3],[4,5,6],[7,8,9]]
// 输出：[1,2,3,6,9,8,7,4,5]
// 示例二
// 输入：matrix = [[1,2,3,4],[5,6,7,8],[9,10,11,12]]
// 输出：[1,2,3,4,8,12,11,10,9,5,6,7]
//
// 提示：
// m == matrix.length
// n == matrix[i].length
// 1 <= m, n <= 10
// -100 <= matrix[i][j] <= 100
//
// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/spiral-matrix
// 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

// spiralOrder
func spiralOrder(matrix [][]int) []int {
	n := len(matrix)
	m := len(matrix[0])
	result := make([]int, 0, m*n)
	top, bottom := 0, n-1 // 高度
	left, right := 0, m-1 // 长度
	for left <= right && top <= bottom {
		// 从左上到右上
		for i := left; i <= right; i++ {
			result = append(result, matrix[top][i])
		}
		// 从右上到右下
		for j := top + 1; j <= bottom; j++ {
			result = append(result, matrix[j][right])
		}
		// 从右下到左下，单数行的话， top 和 bottom相等会出现重复的值
		for i := right - 1; i >= left && top != bottom; i-- {
			result = append(result, matrix[bottom][i])
		}
		// 从左下到左上 同理 单数行，top 和 bottom 相等会出现重复值
		for j := bottom - 1; j > top && left != right; j-- {
			result = append(result, matrix[j][left])
		}
		left++
		right--
		top++
		bottom--
	}
	return result
}
