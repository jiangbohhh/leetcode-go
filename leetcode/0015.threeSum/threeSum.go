package threeSum

import "sort"

// 0015.threeSum
//
// 给你一个包含 n 个整数的数组nums，判断nums中是否存在三个元素 a，b，c ，使得a + b + c = 0 ？请你找出所有和为 0 且不重复的三元组。
//
// 注意：答案中不可以包含重复的三元组。
//
// 输入：nums = [-1,0,1,2,-1,-4]
// 输出：[[-1,-1,2],[-1,0,1]]
//
// 输入：nums = []
// 输出：[]
//
// 输入：nums = [0]
// 输出：[]
//
// 提示：
// 0 <= nums.length <= 3000
// -105 <= nums[i] <= 105
//
// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/3sum
// 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

// 基本思路: 通过将三数之和等于 0 转换成两数之和等于第三数，使用快慢指针来获取两数之和。
// 1先将数组排序，排序后，数组从小到大排列，因此 数字就变成 左侧+中间+右侧的形式，也就是
// ：比如左侧一个数值固定后，通过快慢指针求解中间+右侧，这里需要注意的地方，因为不能重复，
// 所以，每个值都不能取自同一位置，也就是三数索引不会重合。
// 2. 因此，左侧确定后，我们通过遍历 第二层，获取中间的数字，然后三个数相加和 0 对比，如果大
// 于 0，则认为右侧数字过大，将索引向左移动一位，比较是否为0，为 0 则添加进结果集，中间和右侧
// 在相等时停止二层循环，跳出到一层循环。
// 3. 依此，每次重复一层循环后重新计算二层循环，直到结束，


// threeSum
func threeSum(nums []int) [][]int {
	var result = make([][]int, 0)
	sort.Ints(nums)
	n := len(nums)
	for i := 0; i < n; i++ {

		// 保证 第一层不重复
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		end := n - 1
		for start := i + 1; start < n; start++ {
			// 保证第二层不重复
			if start > i + 1 && nums[start] == nums[start - 1] {
				continue
			}

			for start < end && nums[i] + nums[start] + nums[end]  > 0 {
				end--
			}

			// 防止重复，必须使 start < end
			if start == end {
				break
			}

			if nums[i] + nums[start] + nums[end] == 0 {
				result = append(result, []int{nums[i], nums[start], nums[end]})
			}
		}
	}
	return result
}