package twoSum

// 题目： 0001. 两数之和
// 给定一个整数数组 nums 和一个整数目标值 target，请你在该数组中找出和为目标值的那两个整数，并返回它们的数组下标。
// 你可以假设每种输入只会对应一个答案。但是，数组中同一个元素在答案里不能重复出现。
// 你可以按任意顺序返回答案。
//
// 输入：nums = [2,7,11,15], target = 9
// 输出：[0,1]
// 解释：因为 nums[0] + nums[1] == 9 ，返回 [0, 1] 。
//
// 输入：nums = [3,2,4], target = 6
// 输出：[1,2]
//
// 输入：nums = [3,3], target = 6
// 输出：[0,1]
//
// 提示：
// 2 <= nums.length <= 103
// -109 <= nums[i] <= 109
// -109 <= target <= 109
// 只会存在一个有效答案
//
// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/two-sum
// 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
//
//

// 第一种解法：借助哈希表
// 使用哈希表存储 target - i 的值，这样获取 target - i 的时间复杂度就变成 O(1)了。

// 时间复杂度 O(N)，空间复杂度 O(N)

// twoSum
func twoSum(nums []int, target int) []int {
	var m = map[int]int{}
	for i := 0; i < len(nums); i++ {
		if v, ok := m[nums[i]]; ok {
			return []int{v, i}
		} else {
			m[target-nums[i]] = i
		}
	}
	return nil
}
