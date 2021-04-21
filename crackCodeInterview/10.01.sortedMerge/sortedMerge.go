package sortedMerge

import "sort"

// 10.01 合并排序的数组
// 给定两个排序后的数组 A 和 B，其中 A 的末端有足够的缓冲空间容纳 B。 编写一个方法，将 B 合并入 A 并排序。
// 初始化A 和 B 的元素数量分别为m 和 n。
//
// 示例:
// 输入:
// A = [1,2,3,0,0,0], m = 3
// B = [2,5,6],       n = 3
// 输出:[1,2,2,3,5,6]
//
// 说明:
// A.length == n + m
//
// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/sorted-merge-lcci
// 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

// merge
func merge(A []int, m int, B []int, n int) {
	j := 0
	for i := m; i < m+n; i++ {
		A[i] = B[j]
		j++
	}
	sort.Ints(A)
}
