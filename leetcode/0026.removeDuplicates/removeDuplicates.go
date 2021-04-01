package removeDuplicates


// 给你一个有序数组 nums ，请你 原地 删除重复出现的元素，使每个元素 只出现一次 ，返回删除后数组的新长度。
// 不要使用额外的数组空间，你必须在 原地 修改输入数组 并在使用 O(1) 额外空间的条件下完成。
//
// 示例 1：
// 输入：nums = [1,1,2]
// 输出：2, nums = [1,2]
// 解释：函数应该返回新的长度 2 ，并且原数组 nums 的前两个元素被修改为 1, 2 。不需要考虑数组中超出新长度后面的元素。
//
// 示例 2：
// 输入：nums = [0,0,1,1,1,2,2,3,3,4]
// 输出：5, nums = [0,1,2,3,4]
// 解释：函数应该返回新的长度 5 ， 并且原数组 nums 的前五个元素被修改为 0, 1, 2, 3, 4 。不需要考虑数组中超出新长度后面的元素。
//
// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/remove-duplicates-from-sorted-array
// 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

// 第一种解法思路：
// 1. 生成变量 n，在 nums[0:n]上都是无重复元素且升序排列的，初始化时，n = 0。
// 2. 从第二个元素遍历 nums，通过交换 nums[n+1]与后面的元素，如果 nums[i] != nums[n] 表示当前元素可以加入到左侧，
// 交换nums[n+1]和nums[i]的值，并将 n++，如果 nums[i] == nums[n] 说明元素重复了，无需交换，遍历完成就nums[:n]
// 就是有序的。n+1即为有序无重复数组的长度。
// 时间复杂度: O(N)，空间复杂度: O(1)

// removeDuplicates
func removeDuplicates(nums []int) int {
	if len(nums) < 2 {
		return len(nums)
	}
	n := 0
	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[n] {
			nums[n+1], nums[i] = nums[i], nums[n+1]
			n++
		}
	}
	return n+1
}
