package continuousCards

import "sort"

// 剑指 Offer 61. 扑克牌中的顺子
// 从扑克牌中随机抽5张牌，判断是不是一个顺子，即这5张牌是不是连续的。2～10为数字本身，
// A为1，J为11，Q为12，K为13，而大、小王为 0 ，可以看成任意数字。A 不能视为 14。
//
// 示例1:
// 输入: [1,2,3,4,5]
// 输出: True
//
// 示例2:
// 输入: [0,0,1,2,5]
// 输出: True
//
// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/bu-ke-pai-zhong-de-shun-zi-lcof
// 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

// isStraight
func isStraight(nums []int) bool {
	// 排序
	sort.Ints(nums)

	// 记录 0 和相邻牌之间缺少的牌数量
	var zeros, gaps int
	for i := 0; i < len(nums) && nums[i] == 0; i++ {
		zeros++
	}
	left := zeros
	right := left + 1
	for right < len(nums) {
		// 有对子 则不能成为顺子
		if nums[left] == nums[right] {
			return false
		}
		gaps += nums[right] - nums[left] - 1
		left = right
		right++
	}
	if gaps > zeros {
		return false
	}
	return true
}
