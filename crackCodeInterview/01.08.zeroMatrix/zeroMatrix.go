package zeroMatrix

// 01.08: 零矩阵
// 编写一种算法，若M × N矩阵中某个元素为0，则将其所在的行与列清零。
//
// 示例 1：
// 输入：
// [
//  [1,1,1],
//  [1,0,1],
//  [1,1,1]
// ]
// 输出：
// [
//  [1,0,1],
//  [0,0,0],
//  [1,0,1]
// ]
// 示例 2：
// 输入：
// [
//  [0,1,2,0],
//  [3,4,5,2],
//  [1,3,1,5]
// ]
// 输出：
// [
//  [0,0,0,0],
//  [0,4,5,0],
//  [0,3,1,0]
// ]
//
// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/zero-matrix-lcci
// 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

// setZeros 时间复杂度O(m*n) 空间复杂度 O（m+n）
func setZeros(matrix [][]int) {
	// 使用 row 数组的下标记录哪行设置为空，值为 1 则对应下标对应的行去全部为空
	// 使用 col 数组的下标记录哪列设置为空
	row := make([]int, len(matrix))
	col := make([]int, len(matrix[0]))
	for i, v := range matrix {
		for j, v2 := range v {
			if v2 == 0 {
				row[i] = 1
				col[j] = 1
			}
		}
	}
	for i, v := range matrix {
		if row[i] == 1 {
			matrix[i] = make([]int, len(v))
			continue
		}
		for j, _ := range v {
			if col[j] == 1 {
				matrix[i][j] = 0
			}
		}
	}
}
